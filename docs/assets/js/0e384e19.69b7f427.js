"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[976],{1512:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>l,contentTitle:()=>a,default:()=>h,frontMatter:()=>r,metadata:()=>o,toc:()=>c});var i=t(4848),s=t(8453);const r={title:"Introduction",slug:"/"},a=void 0,o={id:"intro",title:"Introduction",description:"AIKit is a comprehensive platform to quickly get started to host, deploy, build and fine-tune large language models (LLMs).",source:"@site/docs/intro.md",sourceDirName:".",slug:"/",permalink:"/aikit/docs/",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/intro.md",tags:[],version:"current",frontMatter:{title:"Introduction",slug:"/"},sidebar:"sidebar",next:{title:"Quick Start",permalink:"/aikit/docs/quick-start"}},l={},c=[{value:"Features",id:"features",level:2}];function d(e){const n={a:"a",code:"code",h2:"h2",li:"li",p:"p",strong:"strong",ul:"ul",...(0,s.R)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(n.p,{children:"AIKit is a comprehensive platform to quickly get started to host, deploy, build and fine-tune large language models (LLMs)."}),"\n",(0,i.jsx)(n.p,{children:"AIKit offers two main capabilities:"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsxs)(n.li,{children:["\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.strong,{children:"Inference"}),": AIKit uses ",(0,i.jsx)(n.a,{href:"https://localai.io/",children:"LocalAI"}),", which supports a wide range of inference capabilities and formats. LocalAI provides a drop-in replacement REST API that is OpenAI API compatible, so you can use any OpenAI API compatible client, such as ",(0,i.jsx)(n.a,{href:"https://github.com/sozercan/kubectl-ai",children:"Kubectl AI"}),", ",(0,i.jsx)(n.a,{href:"https://github.com/sozercan/chatbot-ui",children:"Chatbot-UI"})," and many more, to send requests to open LLMs!"]}),"\n"]}),"\n",(0,i.jsxs)(n.li,{children:["\n",(0,i.jsxs)(n.p,{children:[(0,i.jsx)(n.strong,{children:(0,i.jsx)(n.a,{href:"/aikit/docs/fine-tune",children:"Fine Tuning"})}),": AIKit offers an extensible fine tuning interface. It supports ",(0,i.jsx)(n.a,{href:"https://github.com/unslothai/unsloth",children:"Unsloth"})," for fast, memory efficient, and easy fine-tuning experience."]}),"\n"]}),"\n"]}),"\n",(0,i.jsxs)(n.p,{children:["\ud83d\udc49 To get started, please see ",(0,i.jsx)(n.a,{href:"/aikit/docs/quick-start",children:"Quick Start"}),"!"]}),"\n",(0,i.jsx)(n.h2,{id:"features",children:"Features"}),"\n",(0,i.jsxs)(n.ul,{children:["\n",(0,i.jsx)(n.li,{children:"\ud83d\udca1 No GPU, or Internet access is required for inference!"}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\udc33 No additional tools are needed except for ",(0,i.jsx)(n.a,{href:"https://docs.docker.com/desktop/install/linux-install/",children:"Docker"}),"!"]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83e\udd0f Minimal image size, resulting in less vulnerabilities and smaller attack surface with a custom ",(0,i.jsx)(n.a,{href:"https://github.com/GoogleContainerTools/distroless",children:"distroless"}),"-based image"]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83c\udfb5 ",(0,i.jsx)(n.a,{href:"/aikit/docs/fine-tune",children:"Fine tune support"})]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\ude80 Easy to use declarative configuration for ",(0,i.jsx)(n.a,{href:"/aikit/docs/specs-inference",children:"inference"})," and ",(0,i.jsx)(n.a,{href:"/aikit/docs/specs-finetune",children:"fine tuning"})]}),"\n",(0,i.jsx)(n.li,{children:"\u2728 OpenAI API compatible to use with any OpenAI API compatible client"}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\udcf8 ",(0,i.jsx)(n.a,{href:"/aikit/docs/vision",children:"Multi-modal model support"})]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\uddbc\ufe0f ",(0,i.jsx)(n.a,{href:"/aikit/docs/diffusion",children:"Image generation support"})]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83e\udd99 Support for GGUF (",(0,i.jsx)(n.a,{href:"https://github.com/ggerganov/llama.cpp",children:(0,i.jsx)(n.code,{children:"llama"})}),"), GPTQ or EXL2 (",(0,i.jsx)(n.a,{href:"https://github.com/turboderp/exllamav2",children:(0,i.jsx)(n.code,{children:"exllama2"})}),"), and GGML (",(0,i.jsx)(n.a,{href:"https://github.com/ggerganov/llama.cpp",children:(0,i.jsx)(n.code,{children:"llama-ggml"})}),") and ",(0,i.jsx)(n.a,{href:"https://github.com/state-spaces/mamba",children:"Mamba"})," models"]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\udea2 ",(0,i.jsx)(n.a,{href:"#kubernetes-deployment",children:"Kubernetes deployment ready"})]}),"\n",(0,i.jsx)(n.li,{children:"\ud83d\udce6 Supports multiple models with a single image"}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\udda5\ufe0f Supports ",(0,i.jsx)(n.a,{href:"/aikit/docs/create-images#multi-platform-support",children:"AMD64 and ARM64"})," CPUs and ",(0,i.jsx)(n.a,{href:"/aikit/docs/gpu",children:"GPU-accelerated inferencing with NVIDIA GPUs"})]}),"\n",(0,i.jsxs)(n.li,{children:["\ud83d\udd10 Ensure ",(0,i.jsx)(n.a,{href:"/aikit/docs/security",children:"supply chain security"})," with SBOMs, Provenance attestations, and signed images"]}),"\n",(0,i.jsx)(n.li,{children:"\ud83c\udf08 Support for non-proprietary and self-hosted container registries to store model images"}),"\n"]})]})}function h(e={}){const{wrapper:n}={...(0,s.R)(),...e.components};return n?(0,i.jsx)(n,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},8453:(e,n,t)=>{t.d(n,{R:()=>a,x:()=>o});var i=t(6540);const s={},r=i.createContext(s);function a(e){const n=i.useContext(r);return i.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function o(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:a(e.components),i.createElement(r.Provider,{value:n},e.children)}}}]);