1) docker-compose up --build
2) запустить все контейнеры
3) установить requirements.txt в .venv (pip install -r requirements.txt)
4) запустить файл addDataDB.py (генерация данных)



Список запросов (примеры):
GET: http://localhost:8080/products?lastId=100&pageSize=18

GET: http://localhost:8080/years?model=Sorento

GET: http://localhost:8080/brands

GET: http://localhost:8080/models?brand=Toyota

GET: http://localhost:8080/glassoptions?glasstype=Лобовое стекло

GET: http://localhost:8080/glasstypes

GET: http://localhost:8080/filterproducts?brandName=Toyota&yearModelName=2011-2015&glassTypeName=Лобовое стекло&lastId=0&pageSize=10

POST: http://localhost:8080/products (  price:546
                                        name:tryj
                                        article:rtyj
                                        length:67456
                                        file:
                                        width:4567
                                        amount:4567
                                        brands_name:hgjm
                                        models_name:fgjhgfh
                                        year_model_name:tyu
                                        glass_types_name:fgjh
                                        glass_options_name:jhjh)