generate-mock:
	mockgen -source=interface.go -destination=interface_mockgen.go -package=s3wrapper