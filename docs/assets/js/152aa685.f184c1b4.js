"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[876],{1522:(e,n,t)=>{t.r(n),t.d(n,{assets:()=>c,contentTitle:()=>s,default:()=>u,frontMatter:()=>o,metadata:()=>l,toc:()=>r});var a=t(5893),i=t(1151);const o={title:"Fine Tuning API Specifications"},s=void 0,l={id:"specs-finetune",title:"Fine Tuning API Specifications",description:"v1alpha1",source:"@site/docs/specs-finetune.md",sourceDirName:".",slug:"/specs-finetune",permalink:"/aikit/docs/specs-finetune",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/specs-finetune.md",tags:[],version:"current",frontMatter:{title:"Fine Tuning API Specifications"},sidebar:"sidebar",previous:{title:"Inference API Specifications",permalink:"/aikit/docs/specs-inference"},next:{title:"llama.cpp (GGUF and GGML)",permalink:"/aikit/docs/llama-cpp"}},c={},r=[{value:"v1alpha1",id:"v1alpha1",level:2}];function p(e){const n={code:"code",h2:"h2",p:"p",pre:"pre",...(0,i.a)(),...e.components};return(0,a.jsxs)(a.Fragment,{children:[(0,a.jsx)(n.h2,{id:"v1alpha1",children:"v1alpha1"}),"\n",(0,a.jsx)(n.pre,{children:(0,a.jsx)(n.code,{className:"language-yaml",children:'#syntax=ghcr.io/sozercan/aikit:latest\napiVersion: # required. only v1alpha1 is supported at the moment\nbaseModel: # required. any base model from Huggingface. for unsloth, see for 4bit pre-quantized models: https://huggingface.co/unsloth\ndatasets:\n  - source: # required. this can be a Huggingface dataset repo or a URL pointing to a JSON file\n    type: # required. can be "alpaca". only alpaca is supported at the moment\nconfig:\n  unsloth:\n    packing: # optional. defaults to false. can make training 5x faster for short sequences.\n    maxSeqLength: # optional. defaults to 2048\n    loadIn4bit: # optional. defaults to true\n    batchSize: # optional. default to 2\n    gradientAccumulationSteps: # optional. defaults to 4\n    warmupSteps: # optional. defaults to 10\n    maxSteps: # optional. defaults to 60\n    learningRate: # optional. defaults to 0.0002\n    loggingSteps: # optional. defaults to 1\n    optimizer: # optional. defaults to adamw_8bit\n    weightDecay: # optional. defaults to 0.01\n    lrSchedulerType: # optional. defaults to linear\n    seed: # optional. defaults to 42\noutput:\n  quantize: # optional. defaults to q4_k_m. for unsloth, see for allowed quantization methods: https://github.com/unslothai/unsloth/wiki#saving-to-gguf.\n  name: # optional. defaults to "aikit-model"\n'})}),"\n",(0,a.jsx)(n.p,{children:"Example:"}),"\n",(0,a.jsx)(n.pre,{children:(0,a.jsx)(n.code,{className:"language-yaml",children:"#syntax=ghcr.io/sozercan/aikit:latest\napiVersion: v1alpha1\nbaseModel: unsloth/mistral-7b-instruct-v0.2-bnb-4bit\ndatasets:\n  - source: yahma/alpaca-cleaned\n    type: alpaca\nconfig:\n  unsloth:\n    packing: false\n    maxSeqLength: 2048\n    loadIn4bit: true\n    batchSize: 2\n    gradientAccumulationSteps: 4\n    warmupSteps: 10\n    maxSteps: 60\n    learningRate: 0.0002\n    loggingSteps: 1\n    optimizer: adamw_8bit\n    weightDecay: 0.01\n    lrSchedulerType: linear\n    seed: 42\noutput:\n  quantize: q4_k_m\n  name: model\n"})})]})}function u(e={}){const{wrapper:n}={...(0,i.a)(),...e.components};return n?(0,a.jsx)(n,{...e,children:(0,a.jsx)(p,{...e})}):p(e)}},1151:(e,n,t)=>{t.d(n,{Z:()=>l,a:()=>s});var a=t(7294);const i={},o=a.createContext(i);function s(e){const n=a.useContext(o);return a.useMemo((function(){return"function"==typeof e?e(n):{...n,...e}}),[n,e])}function l(e){let n;return n=e.disableParentContext?"function"==typeof e.components?e.components(i):e.components||i:s(e.components),a.createElement(o.Provider,{value:n},e.children)}}}]);