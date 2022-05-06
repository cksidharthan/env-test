test:
	go test -v env/env_test.go -count 1000


work-test:
	go test -v env/working_env_test.go -count 1000