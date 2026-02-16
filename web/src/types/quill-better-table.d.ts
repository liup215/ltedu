declare module 'quill-table-better' {
  interface QuillTableBetter {
    insertTable(rows: number, columns: number): void;
    deleteTable(): void;
    getTable(range?: any): [any, any, any, number] | null;
    hideTools(): void;
    deleteTableTemporary(source?: any): void;
  }

  interface TableBetterOptions {
    language?: string | { name: string; content: Record<string, any> };
    menus?: string[];
    toolbarTable?: boolean;
    toolbarButtons?: {
      whiteList?: string[];
      singleWhiteList?: string[];
    };
  }

  const QuillTableBetter: {
    new (): QuillTableBetter;
    (quill: any, options: TableBetterOptions): void;
    keyboardBindings: Record<string, any>;
  };

  export default QuillTableBetter;
}