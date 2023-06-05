# building-system-with-the-chatgpt-api

This is the code accompanying the "Building Systems with the ChatGPT API" course by Andrew Ng.
Here is the official website link for the [Building Systems with the ChatGPT API](https://learn.deeplearning.ai/chatgpt-building-system/lesson/1/introduction) course.

lib:
```shell
pip3 install openai
```

Note:
You need to create a .env file in the project directory and write the OpenAI API key in the following format within that file:
```.env
OPENAI_API_KEY ="your-api-key"
```


[get_completion](get_completion.py) This describes how to work with APIs. When using them, you need to remove the "#" symbols to ensure that the code can be executed.

[get_completion_from_messages](get_completion_from_messages.py) Describing how to ensure ChatGPT can output according to predefined rules.

[get_completion_and_token_count](get_completion_and_token_count.py) Describes the calculation rules for tokens when requesting the OpenAI API.