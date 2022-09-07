## Дані і значення
[Розшифровка значень (деякі укр провайдери блокують)](https://rp5.ru/%D0%90%D1%80%D1%85%D1%96%D0%B2_%D0%BF%D0%BE%D0%B3%D0%BE%D0%B4%D0%B8_%D0%B2_%D0%9A%D0%B8%D1%94%D0%B2%D1%96,_%D0%96%D1%83%D0%BB%D1%8F%D0%BD%D0%B0%D1%85_(%D0%B0%D0%B5%D1%80%D0%BE%D0%BF%D0%BE%D1%80%D1%82))

[Класифікація хмар](https://uk.wikipedia.org/wiki/%D0%9C%D1%96%D0%B6%D0%BD%D0%B0%D1%80%D0%BE%D0%B4%D0%BD%D0%B0_%D0%BA%D0%BB%D0%B0%D1%81%D0%B8%D1%84%D1%96%D0%BA%D0%B0%D1%86%D1%96%D1%8F_%D1%85%D0%BC%D0%B0%D1%80)


## Обчислення кількості херових днів по місяцям:
```sql
SELECT count(*), strftime ('%Y-%m',day) as month
FROM (
-- cloudy if there's bad weather clouds
  SELECT AVG(CASE
      WHEN ci OR cm THEN 1
      ELSE 0
  END) as cloudy, strftime ('%Y-%m-%d',timestamp) as day FROM weather_entries GROUP BY strftime ('%Y-%m-%d',timestamp)
)
-- there's 5 data points we consider per day (6:00,9:00,12:00,15:00,18:00), bad weather day if 2 out of 5 were cloudy
WHERE cloudy >=0.4  GROUP BY strftime ('%Y-%m',day)
```

## Results
|Days|Month|
|---|---|
|6|2005-08|
|8|2006-08|
|7|2007-08|
|4|2008-08|
|5|2009-08|
|11|2010-08|
|11|2011-08|
|13|2012-08|
|10|2013-08|
|12|2014-08|
|3|2015-08|
|6|2016-08|
|9|2017-08|
|8|2018-08|
|6|2019-08|
|6|2020-08|
|8|2021-08|
|10|2022-08|

Average without 2022: `7.823529411764706`