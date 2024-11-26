gen/swagger:
	npx redocly bundle ./docs/swagger/root.swagger.yml --output=./docs/swagger/generated.gen.swagger.yml

gen/oapicodegen:
	oapi-codegen -config ./server/config/.codegen.usermanage.config.yml ./docs/swagger/generated.gen.swagger.yml

# make gen/zodios API_NAME=
gen/zodios:
	pnpx openapi-zod-client "./docs/swagger/generated.gen.swagger.yml" -o "./frontend/src/lib/${API_NAME}.ts" --api-client-name ${API_NAME}API
	#sed -i '' '/export function createApiClient/,/}/d' ./frontend/src/lib/${API_NAME}.ts

