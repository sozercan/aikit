"use strict";(self.webpackChunkwebsite=self.webpackChunkwebsite||[]).push([[655],{6716:(e,i,t)=>{t.r(i),t.d(i,{assets:()=>d,contentTitle:()=>a,default:()=>h,frontMatter:()=>s,metadata:()=>r,toc:()=>l});var n=t(5893),o=t(1151);const s={title:"Fine Tuning"},a=void 0,r={id:"fine-tune",title:"Fine Tuning",description:"Fine tuning process allows the adaptation of pre-trained models to domain-specific data. At this time, AIKit fine tuning process is only supported with NVIDIA GPUs.",source:"@site/docs/fine-tune.md",sourceDirName:".",slug:"/fine-tune",permalink:"/aikit/fine-tune",draft:!1,unlisted:!1,editUrl:"https://github.com/sozercan/aikit/blob/main/website/docs/docs/fine-tune.md",tags:[],version:"current",frontMatter:{title:"Fine Tuning"},sidebar:"sidebar",previous:{title:"Creating Model Images",permalink:"/aikit/create-images"},next:{title:"Vision",permalink:"/aikit/vision"}},d={},l=[{value:"Getting Started",id:"getting-started",level:2},{value:"Targets and Configuration",id:"targets-and-configuration",level:2},{value:"Unsloth",id:"unsloth",level:3},{value:"Example Configuration",id:"example-configuration",level:4},{value:"Build",id:"build",level:2},{value:"What&#39;s next?",id:"whats-next",level:2},{value:"Troubleshooting",id:"troubleshooting",level:2},{value:"Build fails with <code>failed to solve: DeadlineExceeded: context deadline exceeded</code>",id:"build-fails-with-failed-to-solve-deadlineexceeded-context-deadline-exceeded",level:3},{value:"Build fails with <code>ERROR 404: Not Found.</code>",id:"build-fails-with-error-404-not-found",level:3}];function c(e){const i={a:"a",admonition:"admonition",code:"code",h2:"h2",h3:"h3",h4:"h4",p:"p",pre:"pre",...(0,o.a)(),...e.components};return(0,n.jsxs)(n.Fragment,{children:[(0,n.jsx)(i.p,{children:"Fine tuning process allows the adaptation of pre-trained models to domain-specific data. At this time, AIKit fine tuning process is only supported with NVIDIA GPUs."}),"\n",(0,n.jsxs)(i.admonition,{type:"note",children:[(0,n.jsx)(i.p,{children:"Due to limitations with BuildKit and NVIDIA, it is essential that the GPU driver version on your host matches the version AIKit will install in the container during the build process."}),(0,n.jsxs)(i.p,{children:["To determine your host GPU driver version, you can execute ",(0,n.jsx)(i.code,{children:"nvidia-smi"})," or ",(0,n.jsx)(i.code,{children:"cat /proc/driver/nvidia/version"}),"."]}),(0,n.jsxs)(i.p,{children:["For information on the GPU driver versions supported by AIKit, please visit ",(0,n.jsx)(i.a,{href:"https://download.nvidia.com/XFree86/Linux-x86_64/",children:"https://download.nvidia.com/XFree86/Linux-x86_64/"}),"."]}),(0,n.jsx)(i.p,{children:"Should your host GPU driver version not be listed, you will need to update to a compatible version available in the NVIDIA downloads mentioned above. It's important to note that there's no need to directly install drivers from the NVIDIA downloads; the versions simply need to be consistent."}),(0,n.jsx)(i.p,{children:"We hope to optimize this process in the future to eliminate this requirement."})]}),"\n",(0,n.jsx)(i.h2,{id:"getting-started",children:"Getting Started"}),"\n",(0,n.jsx)(i.p,{children:"To get started, you need to create a builder to be able to access host GPU devices."}),"\n",(0,n.jsx)(i.p,{children:"Create a builder with the following configuration:"}),"\n",(0,n.jsx)(i.pre,{children:(0,n.jsx)(i.code,{className:"language-bash",children:"docker buildx create --name aikit-builder --use --buildkitd-flags '--allow-insecure-entitlement security.insecure'\n"})}),"\n",(0,n.jsx)(i.admonition,{type:"tip",children:(0,n.jsxs)(i.p,{children:["Additionally, you can build using other BuildKit drivers, such as ",(0,n.jsx)(i.a,{href:"https://docs.docker.com/build/drivers/kubernetes/",children:"Kubernetes driver"})," by setting ",(0,n.jsx)(i.code,{children:"--driver=kubernetes"})," if you are interested in building using a Kubernetes cluster. Please see ",(0,n.jsx)(i.a,{href:"https://docs.docker.com/build/drivers/",children:"BuildKit Drivers"})," for more information."]})}),"\n",(0,n.jsx)(i.h2,{id:"targets-and-configuration",children:"Targets and Configuration"}),"\n",(0,n.jsxs)(i.p,{children:["AIKit is capable of supporting multiple fine tuning implementation targets. At this time, ",(0,n.jsx)(i.a,{href:"https://github.com/unslothai/unsloth",children:"Unsloth"})," is the only supported target, but can be extended for other fine tuning implementations in the future."]}),"\n",(0,n.jsx)(i.h3,{id:"unsloth",children:"Unsloth"}),"\n",(0,n.jsx)(i.p,{children:"Create a YAML file with your configuration. For example, minimum config looks like:"}),"\n",(0,n.jsx)(i.pre,{children:(0,n.jsx)(i.code,{className:"language-yaml",children:'#syntax=ghcr.io/sozercan/aikit:latest\napiVersion: v1alpha1\nbaseModel: "unsloth/llama-2-7b-bnb-4bit" # base model to be fine tuned. this can be any model from Huggingface. For unsloth optimized base models, see https://huggingface.co/unsloth\ndatasets:\n  - source: "yahma/alpaca-cleaned" # data set to be used for fine tuning. This can be a Huggingface dataset or a URL pointing to a JSON file\n    type: "alpaca" # type of dataset. only alpaca is supported at this time.\nconfig:\n  unsloth:\n'})}),"\n",(0,n.jsxs)(i.p,{children:["For full configuration, please refer to ",(0,n.jsx)(i.a,{href:"/aikit/specs-finetune",children:"Fine Tune API Specifications"}),"."]}),"\n",(0,n.jsx)(i.admonition,{type:"note",children:(0,n.jsxs)(i.p,{children:["Please refer to ",(0,n.jsx)(i.a,{href:"https://github.com/unslothai/unsloth",children:"Unsloth documentation"})," for more information about Unsloth configuration."]})}),"\n",(0,n.jsx)(i.h4,{id:"example-configuration",children:"Example Configuration"}),"\n",(0,n.jsx)(i.admonition,{type:"warning",children:(0,n.jsxs)(i.p,{children:["Please make sure to change syntax to ",(0,n.jsx)(i.code,{children:"#syntax=ghcr.io/sozercan/aikit:latest"})," in the example below."]})}),"\n",(0,n.jsx)(i.p,{children:(0,n.jsx)(i.a,{href:"https://github.com/sozercan/aikit/blob/main/test/aikitfile-unsloth.yaml",children:"https://github.com/sozercan/aikit/blob/main/test/aikitfile-unsloth.yaml"})}),"\n",(0,n.jsx)(i.h2,{id:"build",children:"Build"}),"\n",(0,n.jsxs)(i.p,{children:["Build using following command and make sure to replace ",(0,n.jsx)(i.code,{children:"--target"})," with the fine-tuning implementation of your choice (",(0,n.jsx)(i.code,{children:"unsloth"})," is the only option supported at this time), ",(0,n.jsx)(i.code,{children:"--file"})," with the path to your configuration YAML and ",(0,n.jsx)(i.code,{children:"--output"})," with the output directory of the finetuned model."]}),"\n",(0,n.jsx)(i.pre,{children:(0,n.jsx)(i.code,{className:"language-bash",children:'docker buildx build --builder aikit-builder --allow security.insecure --file "/path/to/config.yaml" --output "/path/to/output" --target unsloth --progress plain .\n'})}),"\n",(0,n.jsxs)(i.p,{children:["Depending on your setup and configuration, build process may take some time. At the end of the build, the fine-tuned model will automatically be quantized with the specified format and output to the path specified in the ",(0,n.jsx)(i.code,{children:"--output"}),"."]}),"\n",(0,n.jsxs)(i.p,{children:["Output will be a ",(0,n.jsx)(i.code,{children:"GGUF"})," model file with the name and quanization format from the configuration. For example:"]}),"\n",(0,n.jsx)(i.pre,{children:(0,n.jsx)(i.code,{className:"language-bash",children:"$ ls -al _output\n-rw-r--r--  1 sozercan sozercan 7161089856 Mar  3 00:19 aikit-model-q4_k_m.gguf\n"})}),"\n",(0,n.jsx)(i.h2,{id:"whats-next",children:"What's next?"}),"\n",(0,n.jsxs)(i.p,{children:["\ud83d\udc49 Now that you have a fine-tuned model output as a GGUF file, you can refer to ",(0,n.jsx)(i.a,{href:"/aikit/create-images",children:"Creating Model Images"})," on how to create an image with AIKit to serve your fine-tuned model!"]}),"\n",(0,n.jsx)(i.h2,{id:"troubleshooting",children:"Troubleshooting"}),"\n",(0,n.jsxs)(i.h3,{id:"build-fails-with-failed-to-solve-deadlineexceeded-context-deadline-exceeded",children:["Build fails with ",(0,n.jsx)(i.code,{children:"failed to solve: DeadlineExceeded: context deadline exceeded"})]}),"\n",(0,n.jsxs)(i.p,{children:["This is a known issue with BuildKit and might be related to disk speed. For more information, please see ",(0,n.jsx)(i.a,{href:"https://github.com/moby/buildkit/issues/4327",children:"https://github.com/moby/buildkit/issues/4327"})]}),"\n",(0,n.jsxs)(i.h3,{id:"build-fails-with-error-404-not-found",children:["Build fails with ",(0,n.jsx)(i.code,{children:"ERROR 404: Not Found."})]}),"\n",(0,n.jsx)(i.p,{children:"This issue arises from a discrepancy between the GPU driver versions on your host and the container. Unfortunately, a matching version for your host driver is not available in the NVIDIA downloads at this time. For further details, please consult the note provided at the beginning of this page."}),"\n",(0,n.jsxs)(i.p,{children:["If you are on Windows Subsystem for Linux (WSL), WSL doesn't expose the host driver version information on ",(0,n.jsx)(i.code,{children:"/proc/driver/nvidia/version"}),". Due to this limitation, WSL is not supported at this time."]})]})}function h(e={}){const{wrapper:i}={...(0,o.a)(),...e.components};return i?(0,n.jsx)(i,{...e,children:(0,n.jsx)(c,{...e})}):c(e)}},1151:(e,i,t)=>{t.d(i,{Z:()=>r,a:()=>a});var n=t(7294);const o={},s=n.createContext(o);function a(e){const i=n.useContext(s);return n.useMemo((function(){return"function"==typeof e?e(i):{...i,...e}}),[i,e])}function r(e){let i;return i=e.disableParentContext?"function"==typeof e.components?e.components(o):e.components||o:a(e.components),n.createElement(s.Provider,{value:i},e.children)}}}]);