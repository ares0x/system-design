import os
import openai
import tiktoken
from dotenv import load_dotenv, find_dotenv

_ = load_dotenv(find_dotenv())

openai.api_key = os.environ['OPENAI_API_KEY']

def get_completion(prompt, model='gpt-3.5-turbo'):
    messages = [{"role":"user","content":prompt}]
    response = openai.ChatCompletion.create(
        model = model,
        messages = messages,
        temperature = 0,
    )
    return response.choices[0].message["content"]


#response = get_completion("Waht is the capital of France?")
#print(response)

#response = get_completion("Take the letters in lollipop and reverse them")
#print(response)


#response = get_completion("Take the letters in l-o-l-l-i-p-o-p and reverse them")
#print(response)