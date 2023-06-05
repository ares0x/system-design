import os
import openai
import tiktoken
from dotenv import load_dotenv, find_dotenv

_ = load_dotenv(find_dotenv())

openai.api_key = os.environ['OPENAI_API_KEY']

def get_completion_from_messages(messages, model="gpt-3.5-turbo", temperature=0, max_tokens=500):   
    response = openai.ChatCompletion.create(
        model = model,
        messages = messages,
        temperature = temperature,
        max_tokens = max_tokens,
    ) 
    return response.choices[0].message["content"]

# messages = [  
# {'role':'system', 
#  'content':"""You are an assistant who\
#  responds in the style of Dr Seuss."""},    
# {'role':'user', 
#  'content':"""write me a very short poem\
#  about a happy carrot"""},  
# ] 
# response = get_completion_from_messages(messages, temperature=1)
# print(response)

messages = [
{'role':'system',
 'content':'All your responses must be one sentence long.'},
{'role':'user', 'content':'write me a story about a hannpy carrot'},
]
response =  get_completion_from_messages(messages, temperature=1)
print(response)


