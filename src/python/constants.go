package python

const (
	RequirementsTxt      = "requirements.txt"
	PipfileLock          = "Pipfile.lock"
	Pipfile              = "Pipfile"
	PythonLayer          = "pythonLayer"
	Python3              = "python3"
	Python               = "python"
	DefaultPythonVersion = "3.11.2"
	PythonPath           = "PYTHONPATH"
	BpPythonEnv          = "BP_PYTHON_VERSION"         // Python version specified by buildpack
	BpDefaultPythonEnv   = "BP_DEFAULT_PYTHON_VERSION" // Default Python version
	BpUserPythonEnv      = "BP_USER_PYTHON_VERSION"    // Python version specified by user
)
