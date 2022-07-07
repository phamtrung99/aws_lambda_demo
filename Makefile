sam-run:
	sam build
	sam local start-api -n env.json -p 3001