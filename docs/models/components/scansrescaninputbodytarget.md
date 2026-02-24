# ScansRescanInputBodyTarget


## Supported Types

### One

```go
scansRescanInputBodyTarget := components.CreateScansRescanInputBodyTargetOne(components.One{/* values here */})
```

### Two

```go
scansRescanInputBodyTarget := components.CreateScansRescanInputBodyTargetTwo(components.Two{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scansRescanInputBodyTarget.Type {
	case components.ScansRescanInputBodyTargetTypeOne:
		// scansRescanInputBodyTarget.One is populated
	case components.ScansRescanInputBodyTargetTypeTwo:
		// scansRescanInputBodyTarget.Two is populated
}
```
