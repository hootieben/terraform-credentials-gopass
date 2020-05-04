# terraform-credentials-gopass
A Terraform credentials helper for gopass

An extremely alpha first pass at a terraform credentials helper for gopass. It works on my system (Ubuntu 19.10 eoan) x86_64 and may or may not work on yours

To use run go get and then place the binary in a terraform default search path (generally ~/.terraform.d/plugins). It assumes gopass is already configured and accessible.
Secrets entered should be done in the root store and without folders. Future versions of this will allow passing a folder to prepend to the lookup.
