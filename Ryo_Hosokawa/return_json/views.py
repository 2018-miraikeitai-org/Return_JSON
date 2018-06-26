import json
from django.http import HttpResponse


# Create your views here.
def index(request, data):
    """response を JSON で返却"""
    json_str = json.dumps(data)
    response = HttpResponse(json_str, content_type='application/json; charset=UTF-8')
    return response


def hello_world(request):
    data = {'Result': 'Hello World!'}
    return index(request, data)
