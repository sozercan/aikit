"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[296],{2409:(e,d,r)=>{r.r(d),r.d(d,{assets:()=>t,contentTitle:()=>n,default:()=>o,frontMatter:()=>i,metadata:()=>l,toc:()=>h});var c=r(4848),s=r(8453);const i={title:"Pre-made Models"},n=void 0,l={id:"premade-models",title:"Pre-made Models",description:"AIKit comes with pre-made models that you can use out-of-the-box!",source:"@site/docs/premade-models.md",sourceDirName:".",slug:"/premade-models",permalink:"/aikit/docs/premade-models",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/premade-models.md",tags:[],version:"current",frontMatter:{title:"Pre-made Models"},sidebar:"sidebar",previous:{title:"Quick Start",permalink:"/aikit/docs/quick-start"},next:{title:"Demos",permalink:"/aikit/docs/demo"}},t={},h=[{value:"CPU",id:"cpu",level:2},{value:"NVIDIA CUDA",id:"nvidia-cuda",level:2},{value:"Deprecated Models",id:"deprecated-models",level:2},{value:"CPU",id:"cpu-1",level:3},{value:"NVIDIA CUDA",id:"nvidia-cuda-1",level:3}];function a(e){const d={a:"a",admonition:"admonition",code:"code",h2:"h2",h3:"h3",p:"p",table:"table",tbody:"tbody",td:"td",th:"th",thead:"thead",tr:"tr",...(0,s.R)(),...e.components};return(0,c.jsxs)(c.Fragment,{children:[(0,c.jsx)(d.p,{children:"AIKit comes with pre-made models that you can use out-of-the-box!"}),"\n",(0,c.jsxs)(d.p,{children:["If it doesn't include a specific model, you can always ",(0,c.jsx)(d.a,{href:"https://sozercan.github.io/aikit/premade-models/",children:"create your own images"}),", and host in a container registry of your choice!"]}),"\n",(0,c.jsx)(d.h2,{id:"cpu",children:"CPU"}),"\n",(0,c.jsxs)(d.table,{children:[(0,c.jsx)(d.thead,{children:(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.th,{children:"Model"}),(0,c.jsx)(d.th,{children:"Optimization"}),(0,c.jsx)(d.th,{children:"Parameters"}),(0,c.jsx)(d.th,{children:"Command"}),(0,c.jsx)(d.th,{children:"Model Name"}),(0,c.jsx)(d.th,{children:"License"})]})}),(0,c.jsxs)(d.tbody,{children:[(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3.1"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3.1:8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3.1-8b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://ai.meta.com/llama/license/",children:"Llama"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3.1"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"70B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3.1:70b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3.1-70b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://ai.meta.com/llama/license/",children:"Llama"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\u24c2\ufe0f Mixtral"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"8x7B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/mixtral:8x7b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"mixtral-8x7b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://choosealicense.com/licenses/apache-2.0/",children:"Apache"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83c\udd7f\ufe0f Phi 3.5"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"3.8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/phi3.5:3.8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"phi-3.5-3.8b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://huggingface.co/microsoft/Phi-3.5-mini-instruct/resolve/main/LICENSE",children:"MIT"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83d\udd21 Gemma 2"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"2B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/gemma2:2b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"gemma-2-2b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://ai.google.dev/gemma/terms",children:"Gemma"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\u2328\ufe0f Codestral 0.1"}),(0,c.jsx)(d.td,{children:"Code"}),(0,c.jsx)(d.td,{children:"22B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/codestral:22b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"codestral-22b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://mistral.ai/licenses/MNPL-0.1.md",children:"MNLP"})})]})]})]}),"\n",(0,c.jsx)(d.h2,{id:"nvidia-cuda",children:"NVIDIA CUDA"}),"\n",(0,c.jsxs)(d.table,{children:[(0,c.jsx)(d.thead,{children:(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.th,{children:"Model"}),(0,c.jsx)(d.th,{children:"Optimization"}),(0,c.jsx)(d.th,{children:"Parameters"}),(0,c.jsx)(d.th,{children:"Command"}),(0,c.jsx)(d.th,{children:"Model Name"}),(0,c.jsx)(d.th,{children:"License"})]})}),(0,c.jsxs)(d.tbody,{children:[(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3.1"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama3.1:8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3.1-8b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://ai.meta.com/llama/license/",children:"Llama"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3.1"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"70B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama3.1:70b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3.1-70b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://ai.meta.com/llama/license/",children:"Llama"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\u24c2\ufe0f Mixtral"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"8x7B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/mixtral:8x7b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"mixtral-8x7b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://choosealicense.com/licenses/apache-2.0/",children:"Apache"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83c\udd7f\ufe0f Phi 3.5"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"3.8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/phi3.5:3.8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"phi-3.5-3.8b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://huggingface.co/microsoft/Phi-3.5-mini-instruct/resolve/main/LICENSE",children:"MIT"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83d\udd21 Gemma 2"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"2B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/gemma2:2b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"gemma-2-2b-instruct"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://ai.google.dev/gemma/terms",children:"Gemma"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\u2328\ufe0f Codestral 0.1"}),(0,c.jsx)(d.td,{children:"Code"}),(0,c.jsx)(d.td,{children:"22B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/codestral:22b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"codestral-22b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://mistral.ai/licenses/MNPL-0.1.md",children:"MNLP"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\u2328\ufe0f Flux 1 Dev"}),(0,c.jsx)(d.td,{children:"Text to image"}),(0,c.jsx)(d.td,{children:"12B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/flux1:dev"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"flux-1-dev"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://github.com/black-forest-labs/flux/blob/main/model_licenses/LICENSE-FLUX1-dev",children:"FLUX.1 [dev] Non-Commercial License"})})]})]})]}),"\n",(0,c.jsxs)(d.admonition,{type:"note",children:[(0,c.jsxs)(d.p,{children:["Please see ",(0,c.jsx)(d.a,{href:"https://github.com/sozercan/aikit/tree/main/models",children:"models folder"})," for pre-made model definitions."]}),(0,c.jsx)(d.p,{children:"If not being offloaded to GPU VRAM, minimum of 8GB of RAM is required for 7B models, 16GB of RAM to run 13B models, and 32GB of RAM to run 8x7B models."}),(0,c.jsxs)(d.p,{children:["All pre-made models include CUDA v12 libraries. They are used with ",(0,c.jsx)(d.a,{href:"/aikit/docs/gpu",children:"NVIDIA GPU acceleration"}),". If a supported NVIDIA GPU is not found in your system, AIKit will automatically fallback to CPU with the most optimized runtime (",(0,c.jsx)(d.code,{children:"avx2"}),", ",(0,c.jsx)(d.code,{children:"avx"}),", or ",(0,c.jsx)(d.code,{children:"fallback"}),")."]})]}),"\n",(0,c.jsx)(d.h2,{id:"deprecated-models",children:"Deprecated Models"}),"\n",(0,c.jsx)(d.p,{children:"The following pre-made models are deprecated and no longer updated. Images will continue to be pullable, if needed."}),"\n",(0,c.jsxs)(d.p,{children:["If you need to use these specific models, you can always ",(0,c.jsx)(d.a,{href:"/aikit/docs/create-images",children:"create your own images"}),", and host in a container registry of your choice!"]}),"\n",(0,c.jsx)(d.h3,{id:"cpu-1",children:"CPU"}),"\n",(0,c.jsxs)(d.table,{children:[(0,c.jsx)(d.thead,{children:(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.th,{children:"Model"}),(0,c.jsx)(d.th,{children:"Optimization"}),(0,c.jsx)(d.th,{children:"Parameters"}),(0,c.jsx)(d.th,{children:"Command"}),(0,c.jsx)(d.th,{children:"License"})]})}),(0,c.jsxs)(d.tbody,{children:[(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83d\udc2c Orca 2"}),(0,c.jsx)(d.td,{}),(0,c.jsx)(d.td,{children:"13B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/orca2:13b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE",children:"Microsoft Research"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83c\udd7f\ufe0f Phi 2"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"2.7B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/phi2:2.7b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://huggingface.co/microsoft/phi-2/resolve/main/LICENSE",children:"MIT"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83c\udd7f\ufe0f Phi 3"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"3.8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/phi3:3.8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"phi-3-3.8b"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3:8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3-8b-instruct"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"70B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama3:70b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3-70b-instruct"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 2"}),(0,c.jsx)(d.td,{children:"Chat"}),(0,c.jsx)(d.td,{children:"7B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:7b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-2-7b-chat"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 2"}),(0,c.jsx)(d.td,{children:"Chat"}),(0,c.jsx)(d.td,{children:"13B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/llama2:13b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-2-13b-chat"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83d\udd21 Gemma 1.1"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"2B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm -p 8080:8080 ghcr.io/sozercan/gemma:2b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"gemma-2b-instruct"})})]})]})]}),"\n",(0,c.jsx)(d.h3,{id:"nvidia-cuda-1",children:"NVIDIA CUDA"}),"\n",(0,c.jsxs)(d.table,{children:[(0,c.jsx)(d.thead,{children:(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.th,{children:"Model"}),(0,c.jsx)(d.th,{children:"Optimization"}),(0,c.jsx)(d.th,{children:"Parameters"}),(0,c.jsx)(d.th,{children:"Command"}),(0,c.jsx)(d.th,{children:"License"})]})}),(0,c.jsxs)(d.tbody,{children:[(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83d\udc2c Orca 2"}),(0,c.jsx)(d.td,{}),(0,c.jsx)(d.td,{children:"13B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/orca2:13b-cuda"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://huggingface.co/microsoft/Orca-2-13b/blob/main/LICENSE",children:"Microsoft Research"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83c\udd7f\ufe0f Phi 2"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"2.7B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/phi2:2.7b-cuda"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.a,{href:"https://huggingface.co/microsoft/phi-2/resolve/main/LICENSE",children:"MIT"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83c\udd7f\ufe0f Phi 3"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"3.8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/phi3:3.8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"phi-3-3.8b"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"8B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama3:8b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3-8b-instruct"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 3"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"70B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama3:70b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-3-70b-instruct"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 2"}),(0,c.jsx)(d.td,{children:"Chat"}),(0,c.jsx)(d.td,{children:"7B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama2:7b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-2-7b-chat"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83e\udd99 Llama 2"}),(0,c.jsx)(d.td,{children:"Chat"}),(0,c.jsx)(d.td,{children:"13B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/llama2:13b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"llama-2-13b-chat"})})]}),(0,c.jsxs)(d.tr,{children:[(0,c.jsx)(d.td,{children:"\ud83d\udd21 Gemma 1.1"}),(0,c.jsx)(d.td,{children:"Instruct"}),(0,c.jsx)(d.td,{children:"2B"}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"docker run -d --rm --gpus all -p 8080:8080 ghcr.io/sozercan/gemma:2b"})}),(0,c.jsx)(d.td,{children:(0,c.jsx)(d.code,{children:"gemma-2b-instruct"})})]})]})]})]})}function o(e={}){const{wrapper:d}={...(0,s.R)(),...e.components};return d?(0,c.jsx)(d,{...e,children:(0,c.jsx)(a,{...e})}):a(e)}},8453:(e,d,r)=>{r.d(d,{R:()=>n,x:()=>l});var c=r(6540);const s={},i=c.createContext(s);function n(e){const d=c.useContext(i);return c.useMemo((function(){return"function"==typeof e?e(d):{...d,...e}}),[d,e])}function l(e){let d;return d=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:n(e.components),c.createElement(i.Provider,{value:d},e.children)}}}]);