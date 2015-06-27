dist/: build

build:
	python setup.py sdist bdist

publish: dist/
	twine upload dist/*

clean:
	rm -rf build/ dist/
