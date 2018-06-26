from django.urls import path
from Return_json import views


app_name = 'Return_json'
urlpatterns = [
    path('', views.hello_world, name='hello_world'),
]