package jobsearch

import (
	"log"

	"github.com/tanalam2411/grpc-demo/protogen/basic"
	"github.com/tanalam2411/grpc-demo/protogen/dummy"
	"github.com/tanalam2411/grpc-demo/protogen/jobsearch"
	"google.golang.org/protobuf/encoding/protojson"
)

func JobSearchSoftware() {
	js :=  jobsearch.JobSoftware{
		JobSoftwareId: 777,
		Application: &basic.Application{
			Version: "1.0.0",
			Name: "Job Proto",
			Platforms: []string{"Mac", "Linux", "Windows"},
		},
	}
	
	jsonBytes, _ := protojson.Marshal(&js)
	log.Println("Software: ", string(jsonBytes))
}

func JobSearchCandidate() {
	jc := jobsearch.JobCandidate{
		JobCandidateId: 888,
		Application: &dummy.Application{
			ApplicationId: 887,
			ApplicantFullName: "bam",
			Phone: "12345",
			Email: "bam@bam.com",
		},
	}

	jsonBytes, _ := protojson.Marshal(&jc)
	log.Println("Candidate: ", string(jsonBytes))
}