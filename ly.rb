# typed: false
# frozen_string_literal: true

# This file was generated by GoReleaser. DO NOT EDIT.
class Ly < Formula
  desc ""
  homepage ""
  version "0.0.6"

  on_macos do
    if Hardware::CPU.arm?
      url "https://github.com/jtarchie/ly/releases/download/v0.0.6/ly_darwin_arm64.tar.gz"
      sha256 "f8d3207577c86f690c7802c06232785dbb88023862d7a7e81144ebed4e57348e"

      def install
        bin.install "ly"
      end
    end
    if Hardware::CPU.intel?
      url "https://github.com/jtarchie/ly/releases/download/v0.0.6/ly_darwin_x86_64.tar.gz"
      sha256 "0e69e49f56e278cb8791951d3ed0ba69981d617d225aba5c20bb32e1b2afbbb6"

      def install
        bin.install "ly"
      end
    end
  end

  on_linux do
    if Hardware::CPU.intel?
      url "https://github.com/jtarchie/ly/releases/download/v0.0.6/ly_linux_x86_64.tar.gz"
      sha256 "82ceb0aed4d5fcd3ea36a75104c91c0ddb6d1918182ad5dc36f9a528b3d7c78e"

      def install
        bin.install "ly"
      end
    end
    if Hardware::CPU.arm? && Hardware::CPU.is_64_bit?
      url "https://github.com/jtarchie/ly/releases/download/v0.0.6/ly_linux_arm64.tar.gz"
      sha256 "ecfed49040de2d75d2baa64496ab0598cbea630a6e380aa636a619d172a6255a"

      def install
        bin.install "ly"
      end
    end
  end

  test do
    system "#{bin}/ly --help"
  end
end
