class gno < Formula
    desc "Description of your package"
    homepage "https://github.com/faanrm/go-node"
    url "https://github.com/faanrm/go-node/releases/download/v0.0.1/gno_v0.0.1_linux_386.zip"
    sha256 "sha256_hash_of_your_release_tarball"
  
    def install
      bin.install "gno" # or any other files you need to install
    end
  
    test do
      system "#{bin}/gno", "--version"
    end
  end
  