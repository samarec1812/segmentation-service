segmentation-service
==============================================================================


[![GitHub Actions CI][github-actions-badge]][github-actions-ci-url]

[//]: # ([![Test Coverage][coveralls-badge]][coveralls-badge-url])

[//]: # ([logo]: https://avatars0.githubusercontent.com/u/10262982?v=3&s=150)
[//]: # ([coveralls-badge]: https://img.shields.io/coveralls/samarec1812/currency-rates/master.svg)

[//]: # ([coveralls-badge-url]: https://coveralls.io/github/samarec1812/currency-rates)
[github-actions-badge]: https://github.com/samarec1812/segmentation-service/actions/workflows/go.yml/badge.svg
[github-actions-ci-url]: https://github.com/samarec1812/segmentation-service/actions/workflows/go.yml

Сервис динамического сегментирования пользователей


[//]: # (Features)

[//]: # (------------------------------------------------------------------------------)

[//]: # ()
[//]: # (- Просмотр курса валют относительно указанной даты)

[//]: # (- Выбор валюты по коду валюты из [справочника ЦБР]&#40;https://www.cbr.ru/scripts/XML_val.asp?d=0&#41;)



[//]: # (Installation)

[//]: # (------------------------------------------------------------------------------)

[//]: # ()
[//]: # ()
[//]: # (```)

[//]: # (curl -LO https://github.com/samarec1812/currency-rates/releases/currency-rates_0.0.2_${PLAT}.tar.gz)

[//]: # ()
[//]: # ()
[//]: # (tar -xvf currency-rates_0.0.2_${PLAT}.tar.gz)

[//]: # (```)

[//]: # ()
[//]: # (Usage)

[//]: # (------------------------------------------------------------------------------)

[//]: # (### Precompiled binaries)

[//]: # (```commandline)

[//]: # (currency-rates --code=USD --date=2022-10-08)

[//]: # (```)

[//]: # ()
[//]: # (### Docker images)

[//]: # (```commandline)

[//]: # (docker run --rm samarec1812/currency-rates --code=CNY --date=2022-08-09)

[//]: # (```)

[//]: # (Parameters)

[//]: # (------)

[//]: # ()
[//]: # (| Options | Format &#40;e.g&#41;                                                    | )

[//]: # (|---------|-----------------------------------------------------------------|)

[//]: # (| date    | 2022-10-08                                                      |)

[//]: # (| code    | USD &#40;[all codes]&#40;http://www.cbr.ru/scripts/XML_val.asp?d=0XSD&#41;&#41; |)

[//]: # ()
[//]: # ()
[//]: # (### Comment)

[//]: # (В приложении клиент обращается на адрес:)

[//]: # (https://www.cbr-xml-daily.ru/archive/%s/daily_json.js вместо заданного в исходном API. Это связано с тем, что при обращении клиента на исходный URL сервер отдавал ошибку 403 по независимым для меня причинам. Поэтому было принято решение обращаться на указанный URL.)
