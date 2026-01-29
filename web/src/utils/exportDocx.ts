import { Document, Packer, Paragraph, TextRun, ImageRun } from "docx";
import type { ExamPaper } from "../models/examPaper.model";

/**
 * Export exam paper to docx format
 * @param paper ExamPaper object with name and questions
 */
export async function exportExamPaperToDocx(paper: ExamPaper) {
  const paragraphs: Paragraph[] = [];
  const timestamp = new Date().toISOString().replace(/[-:T]/g, '').slice(0, 14);
  const title = paper.name || 'Exam Paper';

  // Add title
  paragraphs.push(
    new Paragraph({
      children: [
        new TextRun({
          text: title,
          bold: true,
          size: 32,
          font: "Arial"
        }),
        new TextRun({
          text: ' ',
          size: 32,
          font: "Arial"
        })
      ],
      spacing: { after: 400 },
      alignment: "center"
    })
  );

  // Process questions
  const questions = (paper as any).questions || [];
  for (let i = 0; i < questions.length; i++) {
    // Add question number
    paragraphs.push(
      new Paragraph({
        children: [
          new TextRun({
            text: `${i + 1}.`,
            bold: true,
            size: 24,
            font: "Arial"
          }),
          new TextRun({
            text: ' ',
            size: 24,
            font: "Arial"
          })
        ],
        alignment: "left"
      })
    );

    // Process stem content
    if (questions[i].stem) {
      const parser = new DOMParser();
      const doc = parser.parseFromString(questions[i].stem, 'text/html');
      const runs: (TextRun | ImageRun)[] = [];

      // Process text and images
      for (const node of Array.from(doc.body.childNodes)) {
        if (node.nodeType === Node.TEXT_NODE) {
          const text = node.textContent?.trim();
          if (text) {
            runs.push(
              new TextRun({
                text,
                size: 24,
                font: "Arial"
              })
            );
          }
        } else if (node.nodeType === Node.ELEMENT_NODE) {
          const el = node as HTMLElement;
          
          if (el.tagName === 'P') {
            // Process paragraph contents with formatting
            const pRuns: (TextRun | ImageRun)[] = [];
for (const child of Array.from(el.childNodes)) {
  if (child.nodeType === Node.TEXT_NODE) {
    const text = child.textContent?.trim();
    if (text) {
      pRuns.push(
        new TextRun({
          text: text.replace(/\s+/g, ' '), // Normalize whitespace
          size: 24,
          font: "Arial",
          break: text.endsWith('\n') ? 1 : undefined
        })
      );
    }
  } else if (child.nodeType === Node.ELEMENT_NODE) {
    const childEl = child as HTMLElement;
    if (childEl.tagName === 'IMG') {
      try {
        const img = childEl as HTMLImageElement;
        const response = await fetch(img.src);
        const blob = await response.blob();
        const buffer = await blob.arrayBuffer();

        // Calculate image dimensions maintaining aspect ratio
        const imgWidth = img.naturalWidth || 400;
        const imgHeight = img.naturalHeight || 300;
        const maxWidth = 400;
        const maxHeight = 300;
        let finalWidth = imgWidth;
        let finalHeight = imgHeight;

        if (finalWidth > maxWidth) {
          finalHeight = (finalHeight * maxWidth) / finalWidth;
          finalWidth = maxWidth;
        }
        if (finalHeight > maxHeight) {
          finalWidth = (finalWidth * maxHeight) / finalHeight;
          finalHeight = maxHeight;
        }

        pRuns.push(
          new TextRun({ text: "", break: 1 }),
          new ImageRun({
            data: buffer,
            transformation: {
              width: Math.round(finalWidth),
              height: Math.round(finalHeight)
            },
            type: 'png'
          }),
          new TextRun({ text: "", break: 1 })
        );
      } catch {
        pRuns.push(
          new TextRun({
            text: "[图片]",
            size: 24,
            font: "Arial"
          })
        );
      }
    } else if (childEl.tagName === 'SUB') {
      const text = childEl.textContent?.trim();
      if (text) {
        pRuns.push(
          new TextRun({
            text: text.replace(/\s+/g, ' '),
            size: 24,
            font: "Arial",
            subScript: true
          })
        );
      }
    } else if (childEl.tagName === 'SUP') {
      const text = childEl.textContent?.trim();
      if (text) {
        pRuns.push(
          new TextRun({
            text: text.replace(/\s+/g, ' '),
            size: 24,
            font: "Arial",
            superScript: true
          })
        );
      }
    } else if (childEl.classList.contains('ql-formula')) {
      // Try to extract rendered symbol from KaTeX HTML
      const katexHtml = childEl.querySelector('.katex-html');
      let formulaText = '';
      if (katexHtml) {
        formulaText = katexHtml.textContent?.trim() || '';
      } else {
        formulaText = childEl.textContent?.trim() || '';
      }
      if (formulaText) {
        pRuns.push(
          new TextRun({
            text: formulaText,
            size: 24,
            font: "Arial"
          })
        );
      }
    } else {
      const text = childEl.textContent?.trim();
      if (text) {
        // Add the formatted text
        pRuns.push(
          new TextRun({
            text: text.replace(/\s+/g, ' '), // Normalize whitespace
            size: 24,
            font: "Arial",
            bold: childEl.tagName === 'STRONG' || childEl.tagName === 'B',
            underline: childEl.tagName === 'U' ? { type: 'single' } : undefined,
            italics: childEl.tagName === 'I' || childEl.tagName === 'EM',
            break: text.endsWith('\n') ? 1 : undefined
          })
        );
        
        // Add space after bold text
        if (childEl.tagName === 'STRONG' || childEl.tagName === 'B') {
          pRuns.push(
            new TextRun({
              text: ' ',
              size: 24,
              font: "Arial"
            })
          );
        }
      }
    }
  }
}
            
            if (pRuns.length > 0) {
              // Add line break before paragraph if not the first run
              if (runs.length > 0) {
                runs.push(new TextRun({ text: "", break: 1 }));
              }
              
              // Add paragraph content
              runs.push(...pRuns);
              
              // Add line break after paragraph
              runs.push(new TextRun({ text: "", break: 1 }));
            }
          } else if (el.tagName === 'IMG') {
            try {
              const img = el as HTMLImageElement;
              const response = await fetch(img.src);
              const blob = await response.blob();
              const buffer = await blob.arrayBuffer();

              // Add image with breaks before and after
              // Calculate image dimensions maintaining aspect ratio
              const imgWidth = img.naturalWidth || 400;
              const imgHeight = img.naturalHeight || 300;
              const maxWidth = 400;
              const maxHeight = 300;
              let finalWidth = imgWidth;
              let finalHeight = imgHeight;

              if (finalWidth > maxWidth) {
                finalHeight = (finalHeight * maxWidth) / finalWidth;
                finalWidth = maxWidth;
              }
              if (finalHeight > maxHeight) {
                finalWidth = (finalWidth * maxHeight) / finalHeight;
                finalHeight = maxHeight;
              }

              runs.push(
                new TextRun({ text: "", break: 1 }),
                new ImageRun({
                  data: buffer,
                  transformation: {
                    width: Math.round(finalWidth),
                    height: Math.round(finalHeight)
                  },
                  type: 'png'
                }),
                new TextRun({ text: "", break: 1 })
              );
            } catch {
              runs.push(
                new TextRun({
                  text: "[图片]",
                  size: 24,
                  font: "Arial"
                })
              );
            }
          } else if (el.tagName === 'BR') {
            runs.push(new TextRun({ text: "", break: 1 }));
          } else {
            // Process text content in other elements
            Array.from(el.childNodes).forEach(child => {
              if (child.nodeType === Node.TEXT_NODE) {
                const text = child.textContent?.trim();
                if (text) {
                  runs.push(
                    new TextRun({
                      text,
                      size: 24,
                      font: "Arial"
                    })
                  );
                }
              }
            });
          }
        }
      }

      if (runs.length > 0) {
        paragraphs.push(
          new Paragraph({
            children: runs,
            spacing: { after: 200 },
            alignment: "left"
          })
        );
      }
    }
  }

  // Create document with metadata
  const doc = new Document({
    title: title,
    description: `Exam paper export generated on ${new Date().toLocaleString()}`,
    styles: {
      default: {
        document: {
          run: {
            font: "Arial",
            size: 24
          }
        }
      }
    },
    sections: [{
      properties: {
        page: {
          margin: {
            top: 1440,
            right: 1440,
            bottom: 1440,
            left: 1440
          }
        }
      },
      children: paragraphs
    }]
  });

  try {
    const blob = await Packer.toBlob(doc);
    const url = window.URL.createObjectURL(blob);
    const a = document.createElement('a');
    a.href = url;
    a.download = `${title}_${timestamp}.docx`;
    a.click();
    window.URL.revokeObjectURL(url);
  } catch (error) {
    console.error('Failed to generate document:', error);
    throw new Error('Failed to create Word document');
  }
}
