Name:           shop-tui
Version:        0.1.0
Release:        1%{?dist}
Summary:        A Terminal User Interface (TUI) for the Shop-TUI project.

BuildArch:      x86_64

License:        GPL3
URL:            https://github.com/charlie-will-software/shop-tui
Source0:        %{url}/releases/download/v%{version}/%{name}_%{version}_linux_amd64.tar.gz

%description
Your package description.

%prep
tar -xzf %{SOURCE0} -C .

%build
# Nothing to build

%install
rm -rf %{buildroot}
mkdir -p %{buildroot}/%{_bindir}
ls .
cp %{name} %{buildroot}/%{_bindir}

%files
%{_bindir}/%{name}

%changelog
%autochangelog
