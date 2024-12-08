# Shop TUI

## Packaging

### RPM

#### Requirements

```bash
dnf install -y rpmdevtools rpmlint
```

#### Build

```bash
rpmdev-setuptree
cp <REPO_PATH>/tui/shop-tui.spec ./rpmbuild/SPECS/
spectool -gR ~/rpmbuild/SPECS/shop-tui.spec
rpmbuild -ba ~/rpmbuild/SPECS/shop-tui.spec
```

The `.rpm` will then be available under `~/rpmbuild/RPMS/x86_64/`
