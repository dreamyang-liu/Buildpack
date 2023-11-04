yum install -y gzip wget tar
wget https://www.python.org/ftp/python/3.11.2/Python-3.11.2.tgz
tar -zxvf Python-3.11.2.tgz
cd Python-3.11.2
yum install -y xz-devel zlib-devel readline-devel gdb tcl-devel tzdata tk-devel tix-devel openssl-devel autoconf sqlite-devel bluez-libs-devel bzip2 bzip2-devel desktop-file-utils expat-devel findutils gcc-c++ gdbm-devel git-core glibc-all-langpacks glibc-devel gmp-devel libffi-devel libtirpc-devel libGL-devel libuuid-devel libX11-devel make mpdecimal-devel
install_dir=/layers/python-runtime/pythonLayer
./configure --prefix=$install_dir --enable-shared=no CFAGS=-fPIC
make && make install
rm -rf $install_dir/lib/python3.11/test

# Create a tarball of the Python 3.11.2 runtime
cd $install_dir
tar -czvf python3.11.2.tar.gz *
mv python3.11.2.tar.gz /app