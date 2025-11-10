# Project and code guidelines
 - Do not use emojis in any code, scripts or documentation.
 - Use consistent indentation (spaces, not tabs).
 - Prefer clear, concise comments over verbose ones.
 - Follow Go naming conventions (PascalCase for exported identifiers, camelCase for unexported ones).
 - Use descriptive variable and function names.
 - Keep functions small and focused on a single responsibility.
 - Avoid global variables when possible.
 - Use error handling patterns appropriate to the language (e.g., returning errors in Go).
 - Write tests for all new functionality.
 - Document public APIs with godoc comments.
 - Have dedicated files for different functionality (e.g., separate files for models, handlers, utils)
 - Single responsibility principle: each module/class should have one purpose.
 - Ensure code does not expose sensitive information (e.g., secrets, tokens).
 - Take into account OWASP security best practices
 - Do not use ✓ or ✗ in any code, scripts or documentation.
 
# Compiling and Resolving Issues
 - Modify files directly not using sed
 - Read files directly not using sed