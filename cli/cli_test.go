package cli

// func Task() string {
// 	fileContent := `
// ---
// name: test task
// command: do something good
// `
// 	return fileContent
// }
//
// func CreateTestFiles() {
// 	os.MkdirAll("./nginx/tasks", 0700)
// 	fileContent := Task()
// 	ioutil.WriteFile("./nginx/tasks/main.yml", []byte(fileContent), 0700)
// }
//
// func DestroyTestFiles() {
// 	os.RemoveAll("./nginx")
// }
//
// func SetupContext(args string, tag string) *cli.Context {
// 	set := flag.NewFlagSet("test", 0)
// 	set.Parse([]string{args})
// 	set.String("tag", tag, "doc")
// 	return cli.NewContext(nil, set, set)
// }
//
// func TestPushToServer(t *testing.T) {
// 	CreateTestFiles()
// 	defer DestroyTestFiles()
// 	var file multipart.File
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		file, _, _ = r.FormFile("role")
// 	}))
// 	defer ts.Close()
// 	url, _ := url.Parse(ts.URL)
//
// 	tag := url.Host + "/postgres:v1.1"
// 	path := "./nginx"
//
// 	context := SetupContext(path, tag)
//
// 	Push(context)
//
// 	role, err := ioutil.ReadAll(file)
// 	if err != nil {
// 		t.Logf("File not found")
// 		t.FailNow()
// 	}
// 	if len(role) == 0 {
// 		t.Logf("File Empty")
// 		t.FailNow()
// 	}
// 	file.Close()
// 	fileReader := bytes.NewReader(role)
// 	gzipReader, gerror := gzip.NewReader(fileReader)
// 	if gerror != nil {
// 		t.Logf("GZIP ERROR")
// 		t.FailNow()
// 	}
// 	tarBallReader := tar.NewReader(gzipReader)
// 	names := []string{}
// 	for {
// 		header, err := tarBallReader.Next()
// 		if err == io.EOF {
// 			break
// 		}
// 		names = append(names, header.Name)
// 	}
// 	if len(names) != 1 {
// 		t.Fail()
// 	}
// 	if names[0] != "tasks/main.yml" {
// 		t.Log("name not correct " + names[0])
// 		t.Fail()
// 	}
// }
//
// func CreateCompressedFile(file []byte) []byte {
// 	content := []byte{}
// 	tarfile := bytes.NewBuffer(content)
// 	fileWriter := gzip.NewWriter(tarfile)
// 	tarfileWriter := tar.NewWriter(fileWriter)
// 	header := new(tar.Header)
// 	header.Name = "tasks/main.yml"
// 	header.Size = int64(len(file))
// 	header.Mode = int64(0700)
// 	header.Typeflag = tar.TypeReg
// 	tarfileWriter.WriteHeader(header)
// 	tarfileWriter.Write(file)
// 	tarfileWriter.Close()
// 	fileWriter.Close()
// 	return tarfile.Bytes()
// }
//
// func TestPullFromServer(t *testing.T) {
// 	defer os.RemoveAll("./role")
// 	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		file := []byte(Task())
// 		w.Write(CreateCompressedFile(file))
// 	}))
// 	defer ts.Close()
// 	url, _ := url.Parse(ts.URL)
// 	arg := url.Host + "/nginx:v1.1"
// 	context := SetupContext(arg, "")
// 	Clone(context)
// 	if !exists("./role/nginx/tasks/main.yml") {
// 		t.Log("role not found")
// 		t.Fail()
// 	}
// }
