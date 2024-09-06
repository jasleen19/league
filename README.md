# League

## Overview

## Usage

### Run

```
make
````

### Build

```
make build
```

### Test

```
make test
```

### Clean

```
make clean
```

## Client Curl Examples

### Echo

```
curl -F 'file=@/path/to/matrix.csv' "localhost:8080/echo"
```

### Invert

```
curl -F 'file=@/path/to/matrix.csv' "localhost:8080/invert"
```

### Flatten

```
curl -F 'file=@/path/to/matrix.csv' "localhost:8080/flatten"
```

### Sum

```
curl -F 'file=@/path/to/matrix.csv' "localhost:8080/sum"
```

### Multiply

```
curl -F 'file=@/path/to/matrix.csv' "localhost:8080/multiply"
```