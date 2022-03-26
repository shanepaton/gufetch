pkgname=gitfetch
pkgver=1.0.0
pkgrel=1
pkgdesc='Go PKGBUILD Example'
arch=('x86_64')
url="https://github.com/shanepaton/gitfetch"
license=('GPL')
makedepends=('git' 'go')
source=("$pkgname-$pkgver"::'git+https://github.com//shanepaton/gitfetch.git')
sha256sums=('SKIP')

prepare(){
  mkdir -p "$pkgname-$pkgver"
  cd "$pkgname-$pkgver"
  mkdir -p build/
}

build() {
  cd "$pkgname-$pkgver"
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
  go build -o gitfetch
  cp gitfetch build/gitfetch
}

check() {
  cd "$pkgname-$pkgver"
  go test ./...
}

package() {
  cd "$pkgname-$pkgver"
  mkdir -p ~/.config/gitfetch
  [[ -f filename ]] || cp config.yaml ~/.config/gitfetch/
  install -Dm755 build/$pkgname "$pkgdir"/usr/bin/$pkgname
}