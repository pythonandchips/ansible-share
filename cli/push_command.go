package main

func PushRole(path string, walker Walker, compressor Compressor, transport Transport) {
	files := walker.ListFiles(path)
	tarfile := compressor.Compress(path, files)
	transport.UploadFile(tarfile, "role", "role.tar.gz")
}
