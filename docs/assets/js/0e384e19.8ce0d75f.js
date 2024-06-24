"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[671],{7876:(e,t,n)=>{n.r(t),n.d(t,{assets:()=>l,contentTitle:()=>a,default:()=>h,frontMatter:()=>r,metadata:()=>o,toc:()=>c});var i=n(5893),s=n(1151);const r={title:"Introduction",slug:"/"},a=void 0,o={id:"intro",title:"Introduction",description:"AIKit is a comprehensive platform to quickly get started to host, deploy, build and fine-tune large language models (LLMs).",source:"@site/docs/intro.md",sourceDirName:".",slug:"/",permalink:"/aikit/docs/",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/intro.md",tags:[],version:"current",frontMatter:{title:"Introduction",slug:"/"},sidebar:"sidebar",next:{title:"Quick Start",permalink:"/aikit/docs/quick-start"}},l={},c=[{value:"Features",id:"features",level:2}];function d(e){const t={a:"a",code:"code",h2:"h2",li:"li",p:"p",strong:"strong",ul:"ul",...(0,s.a)(),...e.components};return(0,i.jsxs)(i.Fragment,{children:[(0,i.jsx)(t.p,{children:"AIKit is a comprehensive platform to quickly get started to host, deploy, build and fine-tune large language models (LLMs)."}),"\n",(0,i.jsx)(t.p,{children:"AIKit offers two main capabilities:"}),"\n",(0,i.jsxs)(t.ul,{children:["\n",(0,i.jsxs)(t.li,{children:["\n",(0,i.jsxs)(t.p,{children:[(0,i.jsx)(t.strong,{children:"Inference"}),": AIKit uses ",(0,i.jsx)(t.a,{href:"https://localai.io/",children:"LocalAI"}),", which supports a wide range of inference capabilities and formats. LocalAI provides a drop-in replacement REST API that is OpenAI API compatible, so you can use any OpenAI API compatible client, such as ",(0,i.jsx)(t.a,{href:"https://github.com/sozercan/kubectl-ai",children:"Kubectl AI"}),", ",(0,i.jsx)(t.a,{href:"https://github.com/sozercan/chatbot-ui",children:"Chatbot-UI"})," and many more, to send requests to open LLMs!"]}),"\n"]}),"\n",(0,i.jsxs)(t.li,{children:["\n",(0,i.jsxs)(t.p,{children:[(0,i.jsx)(t.strong,{children:(0,i.jsx)(t.a,{href:"/aikit/docs/fine-tune",children:"Fine Tuning"})}),": AIKit offers an extensible fine tuning interface. It supports ",(0,i.jsx)(t.a,{href:"https://github.com/unslothai/unsloth",children:"Unsloth"})," for fast, memory efficient, and easy fine-tuning experience."]}),"\n"]}),"\n"]}),"\n",(0,i.jsxs)(t.p,{children:["\ud83d\udc49 To get started, please see ",(0,i.jsx)(t.a,{href:"/aikit/docs/quick-start",children:"Quick Start"}),"!"]}),"\n",(0,i.jsx)(t.h2,{id:"features",children:"Features"}),"\n",(0,i.jsxs)(t.ul,{children:["\n",(0,i.jsx)(t.li,{children:"\ud83d\udca1 No GPU, or Internet access is required for inference!"}),"\n",(0,i.jsxs)(t.li,{children:["\ud83d\udc33 No additional tools are needed except for ",(0,i.jsx)(t.a,{href:"https://docs.docker.com/desktop/install/linux-install/",children:"Docker"}),"!"]}),"\n",(0,i.jsxs)(t.li,{children:["\ud83e\udd0f Minimal image size, resulting in less vulnerabilities and smaller attack surface with a custom ",(0,i.jsx)(t.a,{href:"https://github.com/GoogleContainerTools/distroless",children:"distroless"}),"-based image"]}),"\n",(0,i.jsxs)(t.li,{children:["\ud83c\udfb5 ",(0,i.jsx)(t.a,{href:"/aikit/docs/fine-tune",children:"Fine tune support"})]}),"\n",(0,i.jsxs)(t.li,{children:["\ud83d\ude80 Easy to use declarative configuration for ",(0,i.jsx)(t.a,{href:"/aikit/docs/specs-inference",children:"inference"})," and ",(0,i.jsx)(t.a,{href:"/aikit/docs/specs-finetune",children:"fine tuning"})]}),"\n",(0,i.jsx)(t.li,{children:"\u2728 OpenAI API compatible to use with any OpenAI API compatible client"}),"\n",(0,i.jsxs)(t.li,{children:["\ud83d\udcf8 ",(0,i.jsx)(t.a,{href:"/aikit/docs/vision",children:"Multi-modal model support"})]}),"\n",(0,i.jsx)(t.li,{children:"\ud83d\uddbc\ufe0f Image generation support with Stable Diffusion"}),"\n",(0,i.jsxs)(t.li,{children:["\ud83e\udd99 Support for GGUF (",(0,i.jsx)(t.a,{href:"https://github.com/ggerganov/llama.cpp",children:(0,i.jsx)(t.code,{children:"llama"})}),"), GPTQ (",(0,i.jsx)(t.a,{href:"https://github.com/turboderp/exllama",children:(0,i.jsx)(t.code,{children:"exllama"})})," or ",(0,i.jsx)(t.a,{href:"https://github.com/turboderp/exllamav2",children:(0,i.jsx)(t.code,{children:"exllama2"})}),"), EXL2 (",(0,i.jsx)(t.a,{href:"https://github.com/turboderp/exllamav2",children:(0,i.jsx)(t.code,{children:"exllama2"})}),"), and GGML (",(0,i.jsx)(t.a,{href:"https://github.com/ggerganov/llama.cpp",children:(0,i.jsx)(t.code,{children:"llama-ggml"})}),") and ",(0,i.jsx)(t.a,{href:"https://github.com/state-spaces/mamba",children:"Mamba"})," models"]}),"\n",(0,i.jsxs)(t.li,{children:["\ud83d\udea2 ",(0,i.jsx)(t.a,{href:"#kubernetes-deployment",children:"Kubernetes deployment ready"})]}),"\n",(0,i.jsx)(t.li,{children:"\ud83d\udce6 Supports multiple models with a single image"}),"\n",(0,i.jsxs)(t.li,{children:["\ud83d\udda5\ufe0f Supports ",(0,i.jsx)(t.a,{href:"/aikit/docs/create-images#multi-platform-support",children:"AMD64 and ARM64"})," CPUs and ",(0,i.jsx)(t.a,{href:"/aikit/docs/gpu",children:"GPU-accelerated inferencing with NVIDIA GPUs"})]}),"\n",(0,i.jsxs)(t.li,{children:["\ud83d\udd10 Ensure ",(0,i.jsx)(t.a,{href:"/aikit/docs/security",children:"supply chain security"})," with SBOMs, Provenance attestations, and signed images"]}),"\n",(0,i.jsx)(t.li,{children:"\ud83c\udf08 Support for non-proprietary and self-hosted container registries to store model images"}),"\n"]})]})}function h(e={}){const{wrapper:t}={...(0,s.a)(),...e.components};return t?(0,i.jsx)(t,{...e,children:(0,i.jsx)(d,{...e})}):d(e)}},1151:(e,t,n)=>{n.d(t,{Z:()=>o,a:()=>a});var i=n(7294);const s={},r=i.createContext(s);function a(e){const t=i.useContext(r);return i.useMemo((function(){return"function"==typeof e?e(t):{...t,...e}}),[t,e])}function o(e){let t;return t=e.disableParentContext?"function"==typeof e.components?e.components(s):e.components||s:a(e.components),i.createElement(r.Provider,{value:t},e.children)}}}]);