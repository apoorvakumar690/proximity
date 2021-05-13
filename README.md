## Directory structure

### apis
It contains controllers for the pkgs that need to be exposed. Each protocol can have different implementations. Each protocol's implementation can reside in its respective directories.

### config
It contains the configuration model. This model will be used in the project.
.yml is used as the configuration file as it is much readable and also supports comments.
The yml config bind to the model which then passed through the project.

### initiate
It would contain code to start up the project. All dependencies would also be created here and then passed to the respective packages.

### pkg
It would mainly contain library code. It could contain multiple packages some of which may depend on other packages. Eventually, some pkg will be used by /apis to get exposed.

### models 
It contains all the database models and all the other models

### services
It contains all the business logic

### vendor
All dependencies would be here.

## Flow
The flow of the program would look like this.

```js
 main.go --> initiate --> config --> apis --> pkg
```
