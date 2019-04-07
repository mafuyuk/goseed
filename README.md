# Goseed
Goseed is a seed command for CSV file.

## Usage
### 1. Getting Goseed
```bash
$ go get github.com/mafuyuk/goseed
```

### 2. Create seeds directory
```bash
$ cd {your repository}
$ mkdir {seeds directory} // e.g. seeds
```

### 3. Create seed file
```bash
$ cd {your seeds directory}
$ vi users.csv
```

#### CSV Hints
File name should be table name
```bash
$ vi users.csv // == users table
```

The first line is the field name
```csv
id,name,description,created_at,updated_at
```

NULL is \N
```csv
id,name,description,created_at,updated_at
1,Jon,\N,1554629710,1554629710
```

### 4. Execute Goseed
```bash
$ goseed -dbname demo -pass pass -user user
```