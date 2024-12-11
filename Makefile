gen/swagger:
	npx redocly bundle ./docs/swagger/root.swagger.yml --output=./docs/swagger/generated.gen.swagger.yml

gen/oapicodegen:
	oapi-codegen -config ./server/config/.codegen.usermanage.config.yml ./docs/swagger/generated.gen.swagger.yml

# make gen/zodios API_NAME=
gen/zodios:
	pnpx openapi-zod-client "./docs/swagger/generated.gen.swagger.yml" -o "./frontend/src/lib/client.ts" --api-client-name usermanageAPI -t "./frontend/config/zodios-template.hbs"

