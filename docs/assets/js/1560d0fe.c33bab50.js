"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[280],{8143:(e,i,n)=>{n.r(i),n.d(i,{assets:()=>o,contentTitle:()=>s,default:()=>p,frontMatter:()=>l,metadata:()=>c,toc:()=>r});var a=n(4848),t=n(8453);const l={title:"llama.cpp (GGUF and GGML)"},s=void 0,c={id:"llama-cpp",title:"llama.cpp (GGUF and GGML)",description:"AIKit utilizes and depends on llama.cpp, which provides inference of Meta's LLaMA model (and others) in pure C/C++, for the llama backend.",source:"@site/docs/llama-cpp.md",sourceDirName:".",slug:"/llama-cpp",permalink:"/aikit/docs/llama-cpp",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/llama-cpp.md",tags:[],version:"current",frontMatter:{title:"llama.cpp (GGUF and GGML)"},sidebar:"sidebar",previous:{title:"Fine Tuning API Specifications",permalink:"/aikit/docs/specs-finetune"},next:{title:"Exllama v2 (GPTQ and EXL2)",permalink:"/aikit/docs/exllama2"}},o={},r=[{value:"Example",id:"example",level:2},{value:"CPU",id:"cpu",level:3},{value:"GPU (CUDA)",id:"gpu-cuda",level:3}];function d(e){const i={a:"a",admonition:"admonition",code:"code",h2:"h2",h3:"h3",li:"li",p:"p",ul:"ul",...(0,t.R)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsxs)(i.p,{children:["AIKit utilizes and depends on ",(0,a.jsx)(i.a,{href:"https://github.com/ggerganov/llama.cpp",children:"llama.cpp"}),", which provides inference of Meta's LLaMA model (and others) in pure C/C++, for the ",(0,a.jsx)(i.code,{children:"llama"})," backend."]}),"\n",(0,a.jsxs)(i.p,{children:["This is the default backend for ",(0,a.jsx)(i.code,{children:"aikit"}),". No additional configuration is required."]}),"\n",(0,a.jsx)(i.p,{children:"This backend:"}),"\n",(0,a.jsxs)(i.ul,{children:["\n",(0,a.jsx)(i.li,{children:"provides support for GGUF (recommended) and GGML models"}),"\n",(0,a.jsxs)(i.li,{children:["supports both CPU (",(0,a.jsx)(i.code,{children:"avx2"}),", ",(0,a.jsx)(i.code,{children:"avx"})," or ",(0,a.jsx)(i.code,{children:"fallback"}),") and CUDA runtimes"]}),"\n"]}),"\n",(0,a.jsx)(i.h2,{id:"example",children:"Example"}),"\n",(0,a.jsx)(i.admonition,{type:"warning",children:(0,a.jsxs)(i.p,{children:["Please make sure to change syntax to ",(0,a.jsx)(i.code,{children:"#syntax=ghcr.io/sozercan/aikit:latest"})," in the examples below."]})}),"\n",(0,a.jsx)(i.h3,{id:"cpu",children:"CPU"}),"\n",(0,a.jsx)(i.p,{children:(0,a.jsx)(i.a,{href:"https://github.com/sozercan/aikit/blob/main/test/aikitfile-llama.yaml",children:"https://github.com/sozercan/aikit/blob/main/test/aikitfile-llama.yaml"})}),"\n",(0,a.jsx)(i.h3,{id:"gpu-cuda",children:"GPU (CUDA)"}),"\n",(0,a.jsx)(i.p,{children:(0,a.jsx)(i.a,{href:"https://github.com/sozercan/aikit/blob/main/test/aikitfile-llama-cuda.yaml",children:"https://github.com/sozercan/aikit/blob/main/test/aikitfile-llama-cuda.yaml"})})]})}function p(e={}){const{wrapper:i}={...(0,t.R)(),...e.components};return i?(0,a.jsx)(i,{...e,children:(0,a.jsx)(d,{...e})}):d(e)}},8453:(e,i,n)=>{n.d(i,{R:()=>s,x:()=>c});var a=n(6540);const t={},l=a.createContext(t);function s(e){const i=a.useContext(l);return a.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function c(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(t):e.components||t:s(e.components),a.createElement(l.Provider,{value:i},e.children)}}}]);