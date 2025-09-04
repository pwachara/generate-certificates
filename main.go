package main

import (
	"bytes"
	_ "embed" // required for go:embed
	"fmt"
	"io"
	"log"
	"net/http"
	"net/mail"
	"strings"

	"github.com/jung-kurt/gofpdf"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/tools/filesystem"
	"github.com/pocketbase/pocketbase/tools/mailer"
)

// -------------------- EMBED ASSETS --------------------
// Put your assets in assets/ before building.
// They are embedded into the binary.
//
//go:embed assets/Allura-Regular.ttf
var embeddedFont []byte

//go:embed assets/certificate-template.jpg
var embeddedTemplate []byte

// -------------------- CONFIG --------------------
const (
	collectionName  = "participants"
	fileFieldName   = "certificate"
	defaultFontSize = 48
	pageWidth       = 297.0 // A4 L width in mm
	pageHeight      = 210.0 // A4 L height in mm
	nameYPos        = 80.0  // tweak to move name up/down
)

// -------------------- MAIN --------------------
func main() {
	app := pocketbase.New()

	// Register routes using pocketbase router
	app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		// 1) Generate all certificates (optionally overwrite with ?force=true)
		se.Router.GET("/api/generate-certificates", func(e *core.RequestEvent) error {
			pbApp, ok := e.App.(*pocketbase.PocketBase)
			if !ok {
				return e.String(http.StatusInternalServerError, "server error: app type assertion failed")
			}

			force := strings.EqualFold(e.Request.URL.Query().Get("force"), "true")

			records, err := pbApp.FindRecordsByFilter(collectionName, "", "", 0, 0)
			if err != nil {
				return e.String(http.StatusInternalServerError, "failed to fetch participants: "+err.Error())
			}

			generated := 0
			for _, rec := range records {
				name := strings.TrimSpace(rec.GetString("name"))
				if name == "" {
					log.Printf("skip %s: empty name\n", rec.Id)
					continue
				}

				// If a certificate already exists and force==false, skip
				if rec.GetString(fileFieldName) != "" && !force {
					continue
				}

				pdfBytes, err := createCertificatePDF(name)
				if err != nil {
					log.Printf("pdf gen error for %s: %v\n", name, err)
					continue
				}

				filename := sanitize(name) + ".pdf"
				fsFile, ferr := filesystem.NewFileFromBytes(pdfBytes, filename)
				if ferr != nil {
					log.Printf("filesystem NewFileFromBytes error for %s: %v\n", name, ferr)
					continue
				}

				rec.Set(fileFieldName, fsFile)
				if err := pbApp.Save(rec); err != nil {
					log.Printf("save record error for %s: %v\n", name, err)
					continue
				}

				log.Printf("generated + saved certificate for %s (record %s)\n", name, rec.Id)
				generated++
			}

			return e.JSON(http.StatusOK, map[string]any{
				"status":    "ok",
				"generated": generated,
			})
		})

		//HTML Email Template to use with the email
				    emailTemplate := `
						<!DOCTYPE html>
						<html>
						<head>
							<meta charset="UTF-8">
							<meta name="viewport" content="width=device-width, initial-scale=1.0">
							<title>AMREF International University Master Class Training Certificate</title>
							<style>
								body {
									font-family: 'Segoe UI', Tahoma, Geneva, Verdana, sans-serif;
									line-height: 1.6;
									color: #333;
									max-width: 600px;
									margin: 0 auto;
									padding: 20px;
									background-color: #f9f9f9;
								}
								.header {
									text-align: center;
									padding: 20px 0;
									background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
									color: white;
									border-radius: 10px 10px 0 0;
								}
								.content {
									background-color: #ffffff;
									padding: 30px;
									border-radius: 0 0 10px 10px;
									box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
								}
								.footer {
									text-align: center;
									margin-top: 30px;
									font-size: 14px;
									color: #666;
								}
								.certificate-icon {
									font-size: 48px;
									text-align: center;
									margin: 20px 0;
									color: #764ba2;
								}
								.button {
									display: inline-block;
									padding: 12px 24px;
									background-color: #764ba2;
									color: white;
									text-decoration: none;
									border-radius: 5px;
									margin: 20px 0;
								}
								.highlight {
									background-color: #f3e8ff;
									padding: 15px;
									border-left: 4px solid #764ba2;
									margin: 20px 0;
								}
							</style>
						</head>
						<body>
							<div class="header">
								<h1>AMREF International University Master Class Training</h1>
								<h3>Certificate of Participation</h3>
							</div>
							
							<div class="content">
								<p>Dear {{NAME}},</p>
								
								<p>Congratulations on successfully completing the AMREF International University Master Class Training held on <strong>29/08/2025</strong>!</p>
								
								<div class="certificate-icon">🎓</div>
								
								<p>We are pleased to provide you with your certificate of participation, which is attached to this email.</p>
								
								<div class="highlight">
									<p>Your certificate is a testament to your dedication and commitment to advancing your knowledge in healthcare.</p>
								</div>
								
								<p>We hope this training has provided you with valuable insights and skills that you can apply in your professional work.</p>
								
								<p>Should you have any questions or need further assistance in regards to this training, please don't hesitate to contact us.</p>
								
								<p>Warm regards,<br>
								<strong>The LOKA ADVISORY Training Team</strong></p>
								
								<div class="footer">
									<p>LOKA ADVISORY SERVICES | Training that equips for impact</p>
									<p>Email: info@loka-advisory.com | Website: www.loka-advisory.com</p>
								</div>
							</div>
						</body>
						</html>
						`
					emailSenderName := "Philemon Wachara"

					emailSubject := "Your AMREF International University Training Certificate"

		// 2) Send all certificates by email (generate+save if missing)
		se.Router.GET("/api/send-certificates", func(e *core.RequestEvent) error {
			pbApp, ok := e.App.(*pocketbase.PocketBase)
			if !ok {
				return e.String(http.StatusInternalServerError, "server error: app type assertion failed")
			}

			records, err := pbApp.FindRecordsByFilter(collectionName, "", "", 0, 0)
			if err != nil {
				return e.String(http.StatusInternalServerError, "failed to fetch participants: "+err.Error())
			}

			mailClient := pbApp.NewMailClient()

			sent := 0
			failed := 0
			for _, rec := range records {
				name := strings.TrimSpace(rec.GetString("name"))
				email := strings.TrimSpace(rec.GetString("email"))
				if name == "" || email == "" {
					log.Printf("skip %s: missing name/email\n", rec.Id)
					failed++
					continue
				}

				pdfBytes, filename, err := getOrCreateCertificate(pbApp, rec, name)
				if err != nil {
					log.Printf("ensure cert error for %s: %v\n", name, err)
					failed++
					continue
				}

				// Build message with attachment (io.Reader)
				msg := &mailer.Message{
					From: mail.Address{Name: emailSenderName, Address: pbApp.Settings().Meta.SenderAddress},
					To:   []mail.Address{{Address: email}},
					Subject: emailSubject,
					HTML:	strings.Replace(emailTemplate, "{{NAME}}", name, 1),
					Attachments: map[string]io.Reader{
						filename: bytes.NewReader(pdfBytes),
					},
				}

				if err := mailClient.Send(msg); err != nil {
					log.Printf("failed to send to %s: %v\n", email, err)
					failed++
					continue
				}

				log.Printf("emailed certificate to %s <%s>\n", name, email)
				sent++
			}

			return e.JSON(http.StatusOK, map[string]any{
				"status": "ok",
				"sent":   sent,
				"failed": failed,
			})
		})

		// 3) Send single certificate by record id (generate+save if missing)
		se.Router.GET("/api/send-certificate/{id}", func(e *core.RequestEvent) error {
			pbApp, ok := e.App.(*pocketbase.PocketBase)
			if !ok {
				return e.String(http.StatusInternalServerError, "server error: app type assertion failed")
			}

			id := e.Request.PathValue("id")
			fmt.Println("This is the selected id ", id)
			if strings.TrimSpace(id) == "" {
				return e.String(http.StatusBadRequest, "missing id parameter")
			}

			rec, err := pbApp.FindRecordById(collectionName, id)
			fmt.Println("This is the record selected from the DB ", rec)
			if err != nil {
				return e.String(http.StatusNotFound, "participant not found")
			}

			name := strings.TrimSpace(rec.GetString("name"))
			email := strings.TrimSpace(rec.GetString("email"))
			if name == "" || email == "" {
				return e.String(http.StatusBadRequest, "participant missing name or email")
			}

			pdfBytes, filename, err := getOrCreateCertificate(pbApp, rec, name)
			if err != nil {
				return e.String(http.StatusInternalServerError, "failed to ensure certificate: "+err.Error())
			}

			mailClient := pbApp.NewMailClient()

			msg := &mailer.Message{
				From: mail.Address{Name: emailSenderName, Address: pbApp.Settings().Meta.SenderAddress},
				To:   []mail.Address{{Address: email}},
				Subject: emailSubject,
				HTML:	strings.Replace(emailTemplate, "{{NAME}}", name, 1),
				Attachments: map[string]io.Reader{
					filename: bytes.NewReader(pdfBytes),
				},
			}

			if err := mailClient.Send(msg); err != nil {
				return e.String(http.StatusInternalServerError, "failed to send email: "+err.Error())
			}

			return e.String(http.StatusOK, "email sent")
		})

		return se.Next()
	})

	// start pocketbase server
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}
}

// -------------------- Helpers --------------------

// getOrCreateCertificate: returns pdf bytes and filename. If record already has a file it tries to read it,
// otherwise generates and saves it into the record's `certificate` file field.
func getOrCreateCertificate(pbApp *pocketbase.PocketBase, rec *core.Record, name string) ([]byte, string, error) {
	// try existing
	existing := strings.TrimSpace(rec.GetString(fileFieldName))
	if existing != "" {
		if b, err := readRecordFile(pbApp, rec, existing); err == nil {
			return b, existing, nil
		}
		// if read failed, we'll regenerate and overwrite
	}

	// generate
	pdfBytes, err := createCertificatePDF(name)
	if err != nil {
		return nil, "", fmt.Errorf("pdf gen: %w", err)
	}
	filename := sanitize(name) + ".pdf"

	fsFile, ferr := filesystem.NewFileFromBytes(pdfBytes, filename)
	if ferr != nil {
		return nil, "", fmt.Errorf("filesystem new file: %w", ferr)
	}

	rec.Set(fileFieldName, fsFile)
	if err := pbApp.Save(rec); err != nil {
		return nil, "", fmt.Errorf("save record: %w", err)
	}

	// final stored name (PocketBase may normalize it)
	finalName := rec.GetString(fileFieldName)
	if finalName == "" {
		finalName = filename
	}
	return pdfBytes, finalName, nil
}

// readRecordFile reads the stored file for a record into memory.
func readRecordFile(pbApp *pocketbase.PocketBase, rec *core.Record, filename string) ([]byte, error) {
	fs, _ := pbApp.NewFilesystem()
	defer fs.Close()

	// rec.BaseFilesPath() is the recommended way to build the path for the record's files
	path := rec.BaseFilesPath() + "/" + filename
	r, err := fs.GetReader(path)
	if err != nil {
		return nil, err
	}
	defer r.Close()
	return io.ReadAll(r)
}

// createCertificatePDF uses embedded font + image and returns the PDF bytes.
func createCertificatePDF(name string) ([]byte, error) {
	pdf := gofpdf.New("L", "mm", "A4", "")
	pdf.AddPage()

	// register font from embedded bytes (no file paths)
	pdf.AddUTF8FontFromBytes("allura", "", embeddedFont)

	// register background image from embedded bytes
	pdf.RegisterImageOptionsReader(
		"template",
		gofpdf.ImageOptions{ImageType: "JPG", ReadDpi: true},
		bytes.NewReader(embeddedTemplate),
	)

	// draw full-page background
	pdf.ImageOptions("template", 0, 0, pageWidth, pageHeight, false, gofpdf.ImageOptions{ImageType: "JPG"}, 0, "")

	// write name
	pdf.SetFont("allura", "", defaultFontSize)
	pdf.SetTextColor(171, 125, 55)
	pdf.SetXY(0, nameYPos)
	pdf.CellFormat(pageWidth, 20, name, "", 0, "C", false, 0, "")

	// output bytes
	var buf bytes.Buffer
	if err := pdf.Output(&buf); err != nil {
		return nil, fmt.Errorf("pdf output failed: %w", err)
	}
	if pdf.Err() {
		return nil, fmt.Errorf("gofpdf internal error: %v", pdf.Error())
	}
	return buf.Bytes(), nil
}

// sanitize filename (simple)
func sanitize(s string) string {
	var b strings.Builder
	for _, r := range s {
		if (r >= 'a' && r <= 'z') ||
			(r >= 'A' && r <= 'Z') ||
			(r >= '0' && r <= '9') {
			b.WriteRune(r)
		} else {
			b.WriteRune('_')
		}
	}
	name := strings.Trim(b.String(), "_")
	if name == "" {
		name = "cert"
	}
	return name
}
