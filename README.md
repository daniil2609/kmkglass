1) docker-compose up --build
2) запустить все контейнеры
3) установить requirements.txt в .venv
4) запустить файл addDataDB.py (генерация данных)



Список запросов (примеры):
http://localhost:8080/products?lastId=100&pageSize=18
http://localhost:8080/years?model=Sorento
http://localhost:8080/brands
http://localhost:8080/models?brand=Toyota
http://localhost:8080/glassoptions?glasstype=Лобовое стекло
http://localhost:8080/glasstypes
http://localhost:8080/filterproducts?brandName=Toyota&yearModelName=2011-2015&glassTypeName=Лобовое стекло&lastId=0&pageSize=10