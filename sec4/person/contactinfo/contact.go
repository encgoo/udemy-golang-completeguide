// subfolder inside subfolder.

// package name is just one layer
package contactinfo

type ContactInfo struct {
	Email   string
	ZipCode int
}

// receiver function
func (pC *ContactInfo) SetEmail(email string) {
	(pC).Email = email
}
