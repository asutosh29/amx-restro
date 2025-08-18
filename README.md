# Restro
## Features
### 1. User
- Place order from the Menu
- Filters and Search are given for ease of access
- User can provide extra instructions (if needed)
- User will be provided with a Table Number and Order ID on placing the order

### 2. Admin
- Access and Manage orders at all stages
- Manage users and give roles (customer and admin)
- Chef page shows all the relevant orders (placed and being cooked) for the Chef
- Pagination for clutter free viewing of all orders

# Docker instructions
If docker is installed
```
> docker compose up --build
```

NOTE: If you are rebuilding the container then make sure to remove the older volume which are created, otherwise the new files won't be copied into the container

# Installation steps
1. Install the required Packages
```
> go mod tidy
```
2. Use the ```.env.sample``` to create your ```.env``` file and fill in the required details

3. Run the following command to get the DB running
```
CREATE DATABASE IF NOT EXIST restro_test
```

Then use the golang migration
```
> migrate -path database/migrations -database "mysql://user:password@tcp(localhost:3306)/restro_test" up
```

4. Start the server
```
> go run cmd/main.go
```


## Usage Instructions
The first user which is created after initialisation is assigned the role ```super``` and has admin access and give or revoke admin privilleges for other users 