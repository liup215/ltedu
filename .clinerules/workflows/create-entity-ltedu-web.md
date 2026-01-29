# Create entity workflow in ltedu-web

## Step 1: Read the entity model and api

- Read the entidy model in `ltedu-api/model/entidy.go` to understand the fields and types.
- Read the API in `ltedu-api/server/api/v1/entidy.go` to understand the endpoints and their parameters.

## Step 2: Create the entity model and Service in ltedu-web

- Create a new file in `ltedu-web/src/models/entidy.ts` for the entity model.
- Create a new file in `ltedu-web/src/services/entidy.ts` for the entity service.

## Step 3: Create Vue pages in ltedu-web


- Create a new directory in `ltedu-web/src/views/admin/EntityManagement.vue` for the entity pages. // if the function is amind management
- Create a new file in `ltedu-web/src/views/admin/EntityForm.vue` for the entity // if the function is admin management

## Step 4: Add routes in ltedu-web

## Step 5: Add the entity to the menu in ltedu-web