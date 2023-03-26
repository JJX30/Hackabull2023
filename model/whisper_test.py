import whisper
import os
import numpy as np
import torch
from IPython.display import Audio

model = whisper.load_model("base")
result = model.transcribe('audio_files/first.mp3', fp16=False)
print(result["text"])