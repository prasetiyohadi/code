# Cloudcustodian

Commands

```bash
# Create Python virtual environment
python -m venv venv
source venv/bin/activate

# Validate policy
custodian validate custodian.yaml

# Dry-run policy
custodian run --dry-run --output-dir=. custodian.yaml
custodian run --dry-run --verbose --output-dir=. custodian.yaml

# Apply policy
custodian run --verbose --output-dir=. custodian.yaml
custodian run --output-dir=. custodian.yaml

# Generate report
custodian report custodian.yaml
custodian report -s . custodian.yaml
custodian report -s . --format grid --field tags=tags custodian.yaml
custodian report --output-dir . --format grid --field tags=tags custodian.yaml
custodian report --output-dir . --format grid --field changed=tags custodian.yaml
custodian report --output-dir . --format grid --field changed=tags.foo custodian.yaml
```
