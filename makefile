generate-mock:
	mockgen -source=interface.go -destination=interface_mockgen.go -package=wsqlx