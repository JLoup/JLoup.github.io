#!/bin/sh

cd resume
go install
cd ..

resume -r me/web/en -baseurl 'https://jloup.github.io/me' > en/JLoupJamet.html
cp en/JLoupJamet.html index.html

resume -r me/en -baseurl /Users/JLoup/dev/jloup.github.io/me > tmp.html
wkhtmltopdf -s A4 --zoom 3.5 tmp.html pdf/en/JLoupJamet.pdf
rm tmp.html

