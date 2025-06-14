import { useMonaco, type Monaco } from "@monaco-editor/react";

import { Editor } from "@monaco-editor/react";
import { useCallback } from "react";
import { Button } from "../ui/button";
import { type editor } from "monaco-editor";
import { Tooltip, TooltipContent, TooltipTrigger } from "../ui/tooltip";

export const CodeEditor = ({
  language,
  defaultValue,
  theme,
  onSubmit,
  isSubmitting,
}: {
  language: string;
  defaultValue: string;
  theme: string;
  onSubmit: (code: string) => void;
  isSubmitting: boolean;
}) => {
  const monaco = useMonaco();

  const handleMount = useCallback(
    (editor: editor.IStandaloneCodeEditor, monaco: Monaco) => {
      setDefaultTypescriptCompilerOptions(monaco);
      addFormatKeybind(editor, monaco);
      addRunKeybind(editor, monaco, onSubmit);
    },
    [onSubmit]
  );

  const handleSubmit = useCallback(() => {
    onSubmit(monaco?.editor.getEditors()[0].getValue() ?? "");
  }, [onSubmit, monaco]);
  console.log(isSubmitting);
  return (
    <div className="w-full h-full relative">
      <Editor
        height="100%"
        defaultLanguage={language}
        defaultValue={defaultValue}
        theme={theme}
        onMount={handleMount}
      />
      <div className="absolute bottom-0 left-0 p-4 flex flex-row gap-2">
        <Tooltip>
          <TooltipTrigger>
            <Button
              variant="default"
              className="cursor-pointer"
              onClick={handleSubmit}
              disabled={isSubmitting}
            >
              {isSubmitting ? "Submitting..." : "Submit"}
            </Button>
          </TooltipTrigger>
          <TooltipContent>
            <p>Ctrl + Enter to submit</p>
          </TooltipContent>
        </Tooltip>
      </div>
    </div>
  );
};
function setDefaultTypescriptCompilerOptions(monaco: Monaco) {
  monaco.languages.typescript.typescriptDefaults.setCompilerOptions({
    target: monaco.languages.typescript.ScriptTarget.ES2020,
    allowNonTsExtensions: true,
    moduleResolution: monaco.languages.typescript.ModuleResolutionKind.NodeJs,
    module: monaco.languages.typescript.ModuleKind.CommonJS,
    lib: ["dom", "dom.iterable", "esnext", "node"],
    types: ["node"],
  });
}

function addFormatKeybind(
  editor: editor.IStandaloneCodeEditor,
  monaco: Monaco
) {
  editor.addAction({
    id: "format",
    label: "Format",
    keybindings: [
      monaco.KeyMod.Shift | monaco.KeyMod.Alt | monaco.KeyCode.KeyF,
    ],
    run: () => {
      editor.getAction("editor.action.formatDocument")?.run();
    },
  });
}

function addRunKeybind(
  editor: editor.IStandaloneCodeEditor,
  monaco: Monaco,
  onSubmit: (code: string) => void
) {
  editor.addAction({
    id: "run",
    label: "Run",
    keybindings: [
      monaco.KeyMod.CtrlCmd | monaco.KeyCode.Enter,
      monaco.KeyMod.WinCtrl | monaco.KeyCode.Enter,
    ],
    run: () => {
      onSubmit(editor.getValue());
    },
  });
}
