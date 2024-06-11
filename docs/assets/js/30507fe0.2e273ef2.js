"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[341],{5772:(e,n,a)=>{a.r(n),a.d(n,{assets:()=>r,contentTitle:()=>l,default:()=>m,frontMatter:()=>o,metadata:()=>s,toc:()=>d});var i=a(5893),t=a(1151);const o={title:"Creating Model Images"},l=void 0,s={id:"create-images",title:"Creating Model Images",description:"This section shows how to create a custom image with models of your choosing. If you want to use one of the pre-made models, skip to running models.",source:"@site/docs/create-images.md",sourceDirName:".",slug:"/create-images",permalink:"/aikit/docs/create-images",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/create-images.md",tags:[],version:"current",frontMatter:{title:"Creating Model Images"},sidebar:"sidebar",previous:{title:"Demos",permalink:"/aikit/docs/demo"},next:{title:"Fine Tuning",permalink:"/aikit/docs/fine-tune"}},r={},d=[{value:"Easy Start",id:"easy-start",level:2},{value:"Build Arguments",id:"build-arguments",level:3},{value:"<code>model</code>",id:"model",level:4},{value:"<code>runtime</code>",id:"runtime",level:4},{value:"Multi-Platform Support",id:"multi-platform-support",level:3},{value:"Advanced Usage",id:"advanced-usage",level:2},{value:"Running models",id:"running-models",level:3}];function c(e){const n={a:"a",admonition:"admonition",code:"code",h2:"h2",h3:"h3",h4:"h4",p:"p",pre:"pre",...(0,t.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.admonition,{type:"note",children:(0,i.jsxs)(n.p,{children:["This section shows how to create a custom image with models of your choosing. If you want to use one of the pre-made models, skip to ",(0,i.jsx)(n.a,{href:"#running-models",children:"running models"}),"."]})}),"\n",(0,i.jsx)(n.p,{children:"First, create a buildx buildkit instance."}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"docker buildx create --use --name aikit-builder\n"})}),"\n",(0,i.jsx)(n.h2,{id:"easy-start",children:"Easy Start"}),"\n",(0,i.jsxs)(n.p,{children:["You can easily build an image from ",(0,i.jsx)(n.a,{href:"https://huggingface.co",children:"Hugging Face"})," models with the following command:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:'docker buildx build -t my-model --load \\\n\t--build-arg="model=huggingface://TheBloke/Llama-2-7B-Chat-GGUF/llama-2-7b-chat.Q4_K_M.gguf" \\\n\t"https://raw.githubusercontent.com/sozercan/aikit/main/models/aikitfile.yaml"\n'})}),"\n",(0,i.jsxs)(n.p,{children:["After building the image, you can proceed to ",(0,i.jsx)(n.a,{href:"#running-models",children:"running models"})," to start the server."]}),"\n",(0,i.jsx)(n.h3,{id:"build-arguments",children:"Build Arguments"}),"\n",(0,i.jsx)(n.p,{children:"Below are the build arguments you can use to customize the image:"}),"\n",(0,i.jsx)(n.h4,{id:"model",children:(0,i.jsx)(n.code,{children:"model"})}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"model"})," build argument is the model URL to download and use. You can use any Hugging Face model URL. Syntax is ",(0,i.jsx)(n.code,{children:"huggingface://foo/bar/baz.gguf"}),". For example:"]}),"\n",(0,i.jsx)(n.p,{children:(0,i.jsx)(n.code,{children:'--build-arg="model=huggingface://TheBloke/Llama-2-7B-Chat-GGUF/llama-2-7b-chat.Q4_K_M.gguf"'})}),"\n",(0,i.jsx)(n.h4,{id:"runtime",children:(0,i.jsx)(n.code,{children:"runtime"})}),"\n",(0,i.jsxs)(n.p,{children:["The ",(0,i.jsx)(n.code,{children:"runtime"})," build argument adds the applicable runtimes to the image. By default, aikit will automatically choose the most optimized CPU runtime. You can use ",(0,i.jsx)(n.code,{children:"cuda"})," to include NVIDIA CUDA runtime libraries. For example:"]}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.code,{children:'--build-arg="runtime=cuda"'}),"."]}),"\n",(0,i.jsx)(n.h3,{id:"multi-platform-support",children:"Multi-Platform Support"}),"\n",(0,i.jsxs)(n.p,{children:["AIKit supports AMD64 and ARM64 multi-platform images. To build a multi-platform image, you can simply add ",(0,i.jsx)(n.code,{children:"--platform linux/amd64,linux/arm64"})," to the build command. For example:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:'docker buildx build -t my-model --load \\\n    --platform linux/amd64,linux/arm64 \\\n    --build-arg="model=huggingface://TheBloke/Llama-2-7B-Chat-GGUF/llama-2-7b-chat.Q4_K_M.gguf" \\\n    "https://raw.githubusercontent.com/sozercan/aikit/main/models/aikitfile.yaml"\n'})}),"\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.a,{href:"https://sozercan.github.io/aikit/docs/premade-models",children:"Pre-made models"})," are offered with multi-platform support. Docker runtime will automatically choose the correct platform to run the image. For more information, please see ",(0,i.jsx)(n.a,{href:"https://docs.docker.com/build/building/multi-platform/",children:"multi-platform images documentation"}),"."]}),"\n",(0,i.jsx)(n.admonition,{type:"note",children:(0,i.jsxs)(n.p,{children:["Please note that ARM64 support only applies to the ",(0,i.jsx)(n.code,{children:"llama.cpp"})," backend with CPU inference. NVIDIA CUDA is not supported on ARM64 at this time."]})}),"\n",(0,i.jsx)(n.h2,{id:"advanced-usage",children:"Advanced Usage"}),"\n",(0,i.jsxs)(n.p,{children:["Create an ",(0,i.jsx)(n.code,{children:"aikitfile.yaml"})," with the following structure:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-yaml",children:"#syntax=ghcr.io/sozercan/aikit:latest\napiVersion: v1alpha1\nmodels:\n  - name: llama-2-7b-chat.Q4_K_M.gguf\n    source: https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGUF/resolve/main/llama-2-7b-chat.Q4_K_M.gguf\n"})}),"\n",(0,i.jsx)(n.admonition,{type:"tip",children:(0,i.jsxs)(n.p,{children:["For full ",(0,i.jsx)(n.code,{children:"aikitfile"})," inference specifications, see ",(0,i.jsx)(n.a,{href:"/aikit/docs/specs-inference",children:"Inference API Specifications"}),"."]})}),"\n",(0,i.jsx)(n.p,{children:"Then build your image with:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"docker buildx build . -t my-model -f aikitfile.yaml --load\n"})}),"\n",(0,i.jsx)(n.p,{children:"This will build a local container image with your model(s). You can see the image with:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"docker images\nREPOSITORY    TAG       IMAGE ID       CREATED             SIZE\nmy-model      latest    e7b7c5a4a2cb   About an hour ago   5.51GB\n"})}),"\n",(0,i.jsx)(n.h3,{id:"running-models",children:"Running models"}),"\n",(0,i.jsx)(n.p,{children:"You can start the inferencing server for your models with:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:'# for pre-made models, replace "my-model" with the image name\ndocker run -d --rm -p 8080:8080 my-model\n'})}),"\n",(0,i.jsxs)(n.p,{children:["You can then send requests to ",(0,i.jsx)(n.code,{children:"localhost:8080"})," to run inference from your models. For example:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:'curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d \'{\n     "model": "llama-2-7b-chat.Q4_K_M.gguf",\n     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]\n   }\'\n'})}),"\n",(0,i.jsx)(n.p,{children:"Output should be similar to:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-json",children:'{\n    "created": 1701236489,\n    "object": "chat.completion",\n    "id": "dd1ff40b-31a7-4418-9e32-42151ab6875a",\n    "model": "llama-2-7b-chat",\n    "choices": [\n        {\n            "index": 0,\n            "finish_reason": "stop",\n            "message": {\n                "role": "assistant",\n                "content": "\\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."\n            }\n        }\n    ],\n    "usage": {\n        "prompt_tokens": 0,\n        "completion_tokens": 0,\n        "total_tokens": 0\n    }\n}\n'})})]})}function m(e={}){const{wrapper:n}={...(0,t.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(c,{...e})}):c(e)}},1151:(e,n,a)=>{a.d(n,{Z:()=>s,a:()=>l});var i=a(7294);const t={},o=i.createContext(t);function l(e){const n=i.useContext(o);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function s(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:l(e.components),i.createElement(o.Provider,{value:n},e.children)}}}]);