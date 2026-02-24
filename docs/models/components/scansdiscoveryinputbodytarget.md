# ScansDiscoveryInputBodyTarget


## Supported Types

### Target1

```go
scansDiscoveryInputBodyTarget := components.CreateScansDiscoveryInputBodyTargetTarget1(components.Target1{/* values here */})
```

### Target2

```go
scansDiscoveryInputBodyTarget := components.CreateScansDiscoveryInputBodyTargetTarget2(components.Target2{/* values here */})
```

## Union Discrimination

Use the `Type` field to determine which variant is active, then access the corresponding field:

```go
switch scansDiscoveryInputBodyTarget.Type {
	case components.ScansDiscoveryInputBodyTargetTypeTarget1:
		// scansDiscoveryInputBodyTarget.Target1 is populated
	case components.ScansDiscoveryInputBodyTargetTypeTarget2:
		// scansDiscoveryInputBodyTarget.Target2 is populated
}
```
