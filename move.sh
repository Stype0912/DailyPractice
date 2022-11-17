for name in `ls *.go`; do
#  mv $name ${name%.*};
  mkdir ${name%.*};
  mv ${name} ./${name%.*};
  done;