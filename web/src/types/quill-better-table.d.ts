declare module 'quill-better-table' {
  interface QuillBetterTable {
    insertTable(rows: number, columns: number): void;
    getTable(): any;
    getTableModule(): any;
  }

  interface OperationMenuOptions {
    items?: {
      insertColumnRight?: { text: string };
      insertColumnLeft?: { text: string };
      insertRowUp?: { text: string };
      insertRowDown?: { text: string };
      mergeCells?: { text: string };
      unmergeCells?: { text: string };
      deleteColumn?: { text: string };
      deleteRow?: { text: string };
      deleteTable?: { text: string };
    };
  }

  interface QuillBetterTableOptions {
    operationMenu?: OperationMenuOptions;
  }

  const QuillBetterTable: {
    new (): QuillBetterTable;
    (quill: any, options: QuillBetterTableOptions): void;
  };

  export default QuillBetterTable;
}