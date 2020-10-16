module goClassroomTools

go 1.15

//Comment these lines out when using for production
//////////////////////////////////////////////////////////
// replace github.com/jdvober/goGoogleAuth => ../goGoogleAuth

replace github.com/jdvober/goClassroomTools/courses => ../goClassroomTools/courses

replace github.com/jdvober/goClassroomTools/courseWork => ../goClassroomTools/courseWork

replace github.com/jdvober/goClassroomTools/students => ../goClassroomTools/students

replace github.com/jdvober/goClassroomTools/studentSubmissions => ../goClassroomTools/studentSubmissions

//////////////////////////////////////////////////////////

require (
	github.com/davecgh/go-spew v1.1.1
	github.com/jdvober/goClassroomTools/courseWork v0.0.0-00010101000000-000000000000
	github.com/jdvober/goClassroomTools/courses v0.0.0-00010101000000-000000000000
	github.com/jdvober/goClassroomTools/studentSubmissions v0.0.0-00010101000000-000000000000
	github.com/jdvober/goClassroomTools/students v0.0.0-00010101000000-000000000000
	github.com/jdvober/goGoogleAuth v0.0.0-20201015184638-396d0d182d1d // indirect
)