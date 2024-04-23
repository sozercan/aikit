"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[341],{5772:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>o,default:()=>m,frontMatter:()=>s,metadata:()=>r,toc:()=>l});var i=t(5893),a=t(1151);const s={title:"Creating Model Images"},o=void 0,r={id:"create-images",title:"Creating Model Images",description:"This section shows how to create a custom image with models of your choosing. If you want to use one of the pre-made models, skip to running models.",source:"@site/docs/create-images.md",sourceDirName:".",slug:"/create-images",permalink:"/aikit/create-images",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/create-images.md",tags:[],version:"current",frontMatter:{title:"Creating Model Images"},sidebar:"sidebar",previous:{title:"Demos",permalink:"/aikit/demo"},next:{title:"Fine Tuning",permalink:"/aikit/fine-tune"}},c={},l=[{value:"Running models",id:"running-models",level:3}];function d(e){const n={a:"a",admonition:"admonition",code:"code",h3:"h3",p:"p",pre:"pre",...(0,a.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.admonition,{type:"note",children:(0,i.jsxs)(n.p,{children:["This section shows how to create a custom image with models of your choosing. If you want to use one of the pre-made models, skip to ",(0,i.jsx)(n.a,{href:"#running-models",children:"running models"}),"."]})}),"\n",(0,i.jsxs)(n.p,{children:["Create an ",(0,i.jsx)(n.code,{children:"aikitfile.yaml"})," with the following structure:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-yaml",children:"#syntax=ghcr.io/sozercan/aikit:latest\napiVersion: v1alpha1\nmodels:\n  - name: llama-2-7b-chat\n    source: https://huggingface.co/TheBloke/Llama-2-7B-Chat-GGUF/resolve/main/llama-2-7b-chat.Q4_K_M.gguf\n"})}),"\n",(0,i.jsx)(n.admonition,{type:"tip",children:(0,i.jsxs)(n.p,{children:["This is the simplest way to get started to build an image. For full ",(0,i.jsx)(n.code,{children:"aikitfile"})," inference specifications, see ",(0,i.jsx)(n.a,{href:"/aikit/specs-inference",children:"Inference API Specifications"}),"."]})}),"\n",(0,i.jsxs)(n.p,{children:["First, create a buildx buildkit instance. Alternatively, if you are using Docker v24 with ",(0,i.jsx)(n.a,{href:"https://docs.docker.com/storage/containerd/",children:"containerd image store"})," enabled, you can skip this step."]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"docker buildx create --use --name aikit-builder\n"})}),"\n",(0,i.jsx)(n.p,{children:"Then build your image with:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"docker buildx build . -t my-model -f aikitfile.yaml --load\n"})}),"\n",(0,i.jsx)(n.p,{children:"This will build a local container image with your model(s). You can see the image with:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:"docker images\nREPOSITORY    TAG       IMAGE ID       CREATED             SIZE\nmy-model      latest    e7b7c5a4a2cb   About an hour ago   5.51GB\n"})}),"\n",(0,i.jsx)(n.h3,{id:"running-models",children:"Running models"}),"\n",(0,i.jsx)(n.p,{children:"You can start the inferencing server for your models with:"}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:'# for pre-made models, replace "my-model" with the image name\ndocker run -d --rm -p 8080:8080 my-model\n'})}),"\n",(0,i.jsxs)(n.p,{children:["You can then send requests to ",(0,i.jsx)(n.code,{children:"localhost:8080"})," to run inference from your models. For example:"]}),"\n",(0,i.jsx)(n.pre,{children:(0,i.jsx)(n.code,{className:"language-bash",children:'curl http://localhost:8080/v1/chat/completions -H "Content-Type: application/json" -d \'{\n     "model": "llama-2-7b-chat",\n     "messages": [{"role": "user", "content": "explain kubernetes in a sentence"}]\n   }\'\n{"created":1701236489,"object":"chat.completion","id":"dd1ff40b-31a7-4418-9e32-42151ab6875a","model":"llama-2-7b-chat","choices":[{"index":0,"finish_reason":"stop","message":{"role":"assistant","content":"\\nKubernetes is a container orchestration system that automates the deployment, scaling, and management of containerized applications in a microservices architecture."}}],"usage":{"prompt_tokens":0,"completion_tokens":0,"total_tokens":0}}\n'})})]})}function m(e={}){const{wrapper:n}={...(0,a.a)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},1151:(e,n,t)=>{t.d(n,{Z:()=>r,a:()=>o});var i=t(7294);const a={},s=i.createContext(a);function o(e){const n=i.useContext(s);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function r(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(a):e.components||a:o(e.components),i.createElement(s.Provider,{value:n},e.children)}}}]);