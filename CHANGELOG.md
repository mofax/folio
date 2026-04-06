# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.6.1] - 2026-04-05

### Added

- **CSS `aspect-ratio`** property on Div elements (CSS Sizing Level 4 §5.1) — derives height from width when no explicit height is set; supports `16 / 9`, `auto 16/9`, single number forms (#112)
- **CSS Color Level 4** space-separated `rgb()`/`hsl()` syntax — `rgb(255 0 0 / 0.5)`, percentage alpha, applies to `rgba()`/`hsla()` too (#108)
- **`html.ParseCSSLength`** public utility — converts CSS length strings (`"1in"`, `"16px"`, `"2em"`, `"50%"`, `calc()`) to PDF points (#109)
- **`Document.ToBytes`** convenience method — returns serialized PDF as `[]byte` for HTTP responses, base64 encoding, in-memory processing (#66)
- **Per-side border-radius on table cells** — `drawCellBordersRounded` now draws each border side independently with corner arcs, instead of requiring all four borders to be identical (#115)
- **WASM header/footer** — `folioRender` accepts `headerHtml`/`footerHtml` in settings JSON, rendered via `SetHeaderElement`/`SetFooterElement` (#102)
- **Invoice example** (`examples/invoice/`) — professional invoice PDF demonstrating rounded table headers, CSS Grid, Flexbox, and optional Tailwind CSS v2

### Fixed

- **Nil font panic** in `runMeasurer` — falls back to Helvetica when `TextRun` has no font set (#98)
- **Table default `border-collapse`** changed from `collapse` to `separate` per CSS 2.1 §17.6 — previously prevented cell border-radius from rendering (#114)
- **Table default margins** removed — browsers set zero margins on `<table>`; added `border-spacing: 2px` default per browser UA stylesheets (#117)
- **Div `drawRoundedBorders`** now uses per-corner radii (`RoundedRectPerCorner`) instead of uniform radius; previously only `r[0]` was used (#104)
- **`:not([hidden])` selector** — attribute selectors inside pseudo-class parens were incorrectly extracted, leaving empty `:not()` that always returned false; enables CSS framework `space-y-*` utilities (#101)
- **`rem` unit parsing** — `parsePlainLength` checked `"em"` suffix before `"rem"`, so `"1rem"` failed to parse (#111)
- **Table cell border-radius in HTML** — converter now skips radius wiring in `border-collapse: collapse` mode per CSS Backgrounds Level 3 §5.3 (#100)
- **README Go version** corrected from 1.21+ to 1.25+ to match go.mod (#124)

### Changed

- **Layout test coverage** 70% → 77.9% — 40 new integration tests for draw functions, table rendering, Div features, Grid layout, Flex column, paragraph indent/ellipsis/orphans
- **Playground URL** updated to `playground.foliopdf.dev`

## [0.6.0] - 2026-04-03

**Breaking changes** — see [MIGRATING.md](MIGRATING.md) for upgrade steps.

### Breaking

- Renamed constructors to `New*`/`Load*`/`Parse*` across `reader`, `barcode`, `layout`, `sign`, `forms`
- `sign.LoadPKCS12` renamed to `sign.ParsePKCS12` (same signature, name-only change)
- `Document.Page(index)` returns `(*Page, error)` instead of panicking
- Unexported internal symbols in `reader` and `svg`
- Baseline positioning uses CSS half-leading with actual font metrics (visual change — text shifts up ~4pt for 12pt Helvetica)
- `vertical-align` accepts length/percentage values (previously ignored)

### Added

- Element-based headers/footers: `SetHeaderElement`, `SetFooterElement`, `SetHeaderText`, `SetFooterText`
- `AddHTMLTemplate` / `AddHTMLTemplateFuncs` for Go template → PDF
- `ValidatePdfA` for early PDF/A validation
- Per-run text highlight (`WithBackgroundColor`, `<mark>` in HTML)
- Inline elements in paragraphs (`<img>`, `<svg>`, `display:inline-block`)
- `<sub>` and `<sup>` rendering with baseline shift and correct spacing
- `baseline-shift` CSS property (keywords and lengths/percentages)
- `vertical-align` extended with length/percentage values per CSS 2.1
- Empty lines from consecutive `\n\n` in paragraphs
- `RunInline` for inline layout elements within paragraphs
- CSS: `text-align-last`, `::marker`, `cmyk()`, `object-fit`, `@supports`, `min()`/`max()`/`clamp()`, `:is()`/`:where()`, repeating gradients, `column-width`/`column-rule`, `string-set`/`string()`, `page-break-inside: avoid`, escape sequences in selectors, multiple `box-shadow`
- WebP and GIF image formats

### Fixed

- `<sub>`/`<sup>` baseline shift — previously only reduced font size (#86)
- Adjacent styled runs no longer insert spurious spaces; inline whitespace collapsing per CSS Text Level 3 §4.1.1 (#86)
- Punctuation after `</sup>`/`</sub>` keeps correct styling (#86)
- Punctuation at font boundaries keeps its own font (#30)
- Consecutive `\n\n` produce visible empty lines (#91)
- Blank lines preserved across page splits (#95)
- Paragraph baseline uses CSS half-leading `(lineH + ascent - descent) / 2` (#90)
- `cloneWithWords` preserves all Word styling + line breaks on page-split paragraphs
- Style changes at line break boundaries preserved during page splits
- Overflow handling includes following siblings in Div layout (#13)
- Table layout handles zero/negative height without panicking
- Inline-block SVG/IMG dispatch to correct converters (#71)
- Inline element alignment: line-relative child positions (#71)
- `buildParagraphFromRuns` always uses `NewStyledParagraph` (was dropping `BaselineShift`/`BackgroundColor`)
- RGB color components clamped to 0–1
- Case-insensitive attribute selector matching
- Alpha premultiplication fix for PNG
- Font descriptor flags from actual metadata
- Kern format 0 nPairs validated
- Encrypted PDFs detected with clear error
- Signatures preserved on multi-sign PDFs
- Highlight/underline/strikethrough use actual font metrics (#73)
- Predictor column count bounded to prevent allocation DoS

### Contributors

- **Ben Davidson** ([@bendavidsonku](https://github.com/bendavidsonku)) — inline elements in paragraphs, per-run text highlight background (#71, #72)
- **Jason Kulatunga** ([@AnalogJ](https://github.com/AnalogJ)) — table zero-height fix, overflow sibling handling (#13)

### Changed

- `golang.org/x/image` v0.37.0 → v0.38.0
- Internal: `html/converter.go` split into focused modules
- Internal: `ARCHITECTURE.md` with design principles, layering rules, naming conventions

## [0.5.2] - 2026-03-26

### Added
- **PDF redaction** — `RedactText`, `RedactPattern`, `RedactRegions` permanently remove text from content streams with character-level TJ splitting precision; configurable fill color, overlay text, and metadata stripping (#59)
- **Page import** — `Page.ImportPage` and `Page.ImportPageWithOpts` load existing PDF pages as Form XObjects (ISO 32000 §8.10) for template workflows; `reader.ExtractPageImport` convenience API with full indirect-ref resolution (#47)
- **Drawing primitives** — `Page.AddLine`, `Page.AddRect`, `Page.AddRectFilled` for low-level graphics on pages
- **PDF/UA accessibility** — alt text for images, custom structure tags, structure tree reading from existing PDFs (#60)
- **Paragraph `\n` line breaks** — `\n`, `\r\n`, and `\r` now produce forced line breaks in paragraphs and table cells (#61, #63)
- **C ABI expanded to 346 functions** (up from 330) — adds redaction, page import, drawing, digital signatures, encryption permissions, page manipulation, content extraction, form flattening, merge, TextRun builder, styled list/heading exports
- **Examples** — `merge/` (parse, merge, extract text), `sign/` (PAdES B-B digital signature), `report/` (multi-page layout API), `import-page/` (external PDF template filling), `redact/` (sensitive text removal)

### Fixed
- **Import page blank output** — `resolveDeep` recursively resolves all indirect references in imported resources; `hoistStreams` converts nested PdfStream objects to indirect refs, fixing blank/partial output with real-world PDFs
- **Paragraph `\n` collapsed to space** — `splitWords` now splits on newlines first and inserts forced line break markers

## [0.5.1] - 2026-03-25

### Fixed
- **Release workflow** — replaced deprecated `macos-13` runner with `macos-latest` for x86_64 builds
- **Fuzz test regex** — anchored `-fuzz='^FuzzParse$'` to avoid matching `FuzzParsePDF`

## [0.5.0] - 2026-03-25

### Contributors

- **Marc Ole Bulling** — `<br>` nil pointer fix (#10)
- **Moritz** ([@FrauElster](https://github.com/FrauElster)) — PDF/A-3b file attachments (#17)
- **Piotr Pawlak** ([@piotrxp](https://github.com/piotrxp)) — SSRF prevention for remote resources (#39)

### Added
- **C ABI expanded to 281 functions** (up from 115) — covers nearly all Go engine features
- **Barcode C ABI** — `folio_barcode_qr`, `qr_ecc`, `code128`, `ean13` + layout elements
- **SVG C ABI** — `folio_svg_parse`, `parse_bytes` + layout elements with size/align
- **Link C ABI** — hyperlink, embedded font, and internal link layout elements
- **Flex C ABI** — full flexbox container with items, direction, justify, align, wrap, gap, borders
- **Grid C ABI** — CSS Grid with template columns/rows, auto-rows, placement, justify/align items/content
- **Columns C ABI** — multi-column layout with gap and custom widths
- **Float C ABI** — left/right floating elements with margin
- **TabbedLine C ABI** — tab-stop text with dot leaders for TOC-style layouts
- **Form filling C ABI** — `folio_form_filler_new`, `set_value`, `set_checkbox`, `field_names`, `get_value`
- **Form field builder C ABI** — `folio_form_create_text_field`, `create_checkbox` + `set_value`, `set_read_only`, `set_required`, `set_background_color`, `set_border_color`, then `add_field`
- **Additional form fields** — multiline text, password, listbox, radio group
- **Document watermark** — `folio_document_set_watermark` and `set_watermark_config`
- **Outlines/bookmarks C ABI** — `folio_document_add_outline`, `add_outline_xyz`, `outline_add_child`
- **Named destinations** — `folio_document_add_named_dest`
- **Viewer preferences** — `folio_document_set_viewer_preferences`
- **Page labels** — `folio_document_add_page_label`
- **File attachments** — `folio_document_attach_file` for PDF/A-3b compliance
- **Inline HTML** — `folio_document_add_html` and `add_html_with_options`
- **Page-specific margins** — `folio_document_set_first_margins`, `set_left_margins`, `set_right_margins`
- **Absolute positioning** — `folio_document_add_absolute`
- **Page extensions** — art box, page size override, page-to-page links, text annotations, text markup annotations (highlight, underline, squiggly, strikeout), separate fill/stroke opacity
- **All 14 standard font accessors** — added `helvetica_oblique`, `helvetica_bold_oblique`, `times_italic`, `times_bold_italic`, `courier_bold`, `courier_oblique`, `courier_bold_oblique`, `symbol`, `zapf_dingbats`
- **Paragraph extensions** — orphans, widows, ellipsis, word-break, hyphens
- **Table extensions** — footer rows, cell spacing, auto column widths, min width
- **Cell extensions** — per-side padding, vertical alignment, borders, width hints
- **Div extensions** — border radius, opacity, overflow, max/min width, box shadow, max height
- **List extensions** — leading, nested sub-lists
- **Image extensions** — TIFF loading, element alignment
- **C ABI audit script** (`scripts/audit-cabi.sh`) — detects drift between Go exports, folio.h, and built symbols
- **C integration tests** — 258 tests covering all exported functions

### Changed
- **Library version injected at build time** via `-ldflags "-X main.version=..."` — `folio_version()` returns the git tag in releases, `git describe` in dev builds
- **CI** — added C ABI audit, shared library build, and C integration test steps
- **Release workflow** — builds native shared libraries for 5 platforms (linux-x86_64, linux-aarch64, macos-x86_64, macos-aarch64, windows-x86_64) with SHA256 checksums
- **Makefile** — added `audit`, `audit-build`, cross-compilation targets (`cross-linux-amd64`, etc.), OS-aware shared library extension detection

### Fixed
- **`<br>` nil pointer** — fixed nil pointer exception with `<br>` tags (#10)
- **PDF/A-3b file attachments** — proper embedded file streams with `/AF`, `/Names`, MIME types (#17)
- **ZUGFeRD lint** — deterministic timestamps, fixed example (#41)
- **Content stream compression** — FlateDecode for content streams, merge stream dict bug (#44)
- **SSRF prevention** — URL policy interceptor for blocking/modifying remote resource requests (#39)
- **Radio button / checkbox appearance** — fixed appearance stream generation for form fields
- **`:root` selector** — now correctly matches the `<html>` element
- **Gradient rendering** — fixed CSS gradient parsing and rendering
- **Page number CSS counter** — `counter(page)` and `counter(pages)` now work in margin boxes
- **MarginBox API** — exposed `MarginBoxes`/`FirstMarginBoxes` on `ConvertResult` for simpler programmatic access (#54)

## [0.4.2] - 2026-03-22

### Added
- **Clickable links in PDFs** — `<a href="...">` inside paragraphs, headings, and list items now produce PDF link annotations (#23, #26, #27)
- **Multiple links per line** — paragraphs with several inline links each get their own precise annotation rectangle
- **Internal document links** — `layout.NewInternalLink` resolves to direct page references for macOS Preview compatibility
- **Layout API link support** — `TextRun.WithLinkURI()` and `WithDecoration()` for building linked text programmatically; `List.AddItemRuns()` for linked list items
- **Links example** (`examples/links/`) showcasing external, inline, multi-line, styled, heading, and list item links plus bookmarks and internal navigation
- **Fonts example** (`examples/fonts/`) demonstrating custom `@font-face` with Unicode (CJK, Cyrillic, Japanese)

### Fixed
- **Custom `@font-face` family names ignored** — `parseFontFamily` was mapping all names to standard fonts; now preserves custom names for embedded font matching (#16)
- **`page-break-after` ignored when body has `width: 100%`** — `AreaBreak` elements trapped inside Div wrappers are now hoisted out so the renderer can act on them (#21)
- **CSS class selectors case-insensitive** — `.myClass` now matches `class="myClass"` regardless of case (#28)
- **Punctuation spacing at run boundaries** — period/comma after a styled span (e.g. `<b>word</b>.`) no longer gets an extra inter-word space (#25)
- **Underline continuous across multi-word links** — decoration extends through trailing spaces between consecutive linked/decorated words
- **`@font-face` family name case mismatch** — font-face names are now lowercased consistently so CSS lookup matches

### Changed
- **Split `html/converter.go`** into 11 focused files by responsibility (paragraph, table, block, flex, forms, image, list, heading, link, style, helpers) — no behavior changes (#34)

## [0.4.1] - 2026-03-22

### Contributors

- **Emrecan BATI** — Apache 2.0 license text cleanup

### Added
- **Comprehensive GoDoc comments** across all 14 packages — every exported and unexported symbol now has an accurate doc comment following Go conventions
- **Package-level doc comments** added to `layout`, `html`, `svg`, and consolidated in `core` (had two conflicting comments)
- **`ARCHITECTURE.md`** documenting design principles, package responsibilities, layering rules, dependency policy, and non-goals
- **Examples directory** with hello-world sample

### Fixed
- Stale/inaccurate doc comments: watermark "prepends" → "appends", `dss.Build` return description, `PdfObjects` → `BuildObjects`, merged interface docs in layout, and others
- `cmd/folio printUsage` referenced nonexistent "region" extraction strategy
- `svg/doc.go` listed nonexistent `clipPath` support, missing `radialGradient`
- `layout/doc.go` referenced wrong type names (`Tab` → `TabbedLine`, nonexistent `Transform` element)
- `font/standard.go` package doc was stale ("and later font parsing" — parsing is fully implemented)

### Changed
- Removed committed `folio.wasm` binary (7.2MB) from the repository; the release workflow already builds it fresh per tagged version
- Added `*.wasm` to `.gitignore`
- golangci-lint issues fixed and linter added to CI
- Apache 2.0 license replaced with verbatim text; missing license headers added

## [0.4.0] - 2026-03-19

### Added
- **WOFF1 font decoding** — `@font-face` now supports `.woff` files via automatic format detection (`font.LoadFont`)
- **CSS custom properties (variables)** — `--name: value` declarations with `var(--name, fallback)` resolution, inheritance, and nesting
- **CSS counters** — `counter-reset`, `counter-increment`, `counter()` and `counters()` in `::before`/`::after` content
- **CSS `clear` property** — `clear: left/right/both` advances past active floats before placing elements
- **CSS `border-spacing`** — horizontal and vertical cell spacing for tables in separate border model
- **HTTP background images** — `background-image: url('https://...')` fetches remote images in non-WASM builds
- **Inline-block in text flow** — `display: inline-block` elements flow within paragraphs as "big words" with correct line-breaking and height expansion
- **Containing-block absolute positioning** — `position: absolute` resolves against nearest positioned ancestor via overlay children, not just the page
- **Liang-Knuth hyphenation** — 4938 TeX US English patterns for linguistically correct syllable breaks, replacing geometric character splitting
- **C ABI export layer** — `export/` package for FFI from Python, Ruby, Swift, etc.
- **Full sRGB ICC profile** and PDF/A-1b compliance support
- **QR code v1-40** with numeric/alphanumeric encoding modes
- **Symbol and ZapfDingbats** font width tables

### Changed
- `font.LoadTTF` calls in HTML converter replaced with `font.LoadFont` (auto-detects TTF/OTF/WOFF)
- `hyphenateWord()` uses pattern-based breaks first, falls back to character splitting

## [0.3.0] - 2026-03-18

### Added
- **CSS Grid layout** — `display: grid` with `grid-template-columns`, `grid-template-rows`, `grid-template-areas`, named areas, `auto-rows`, alignment (`justify-items`, `align-items`), and page break support
- **Absolute positioning with z-index** — `position: absolute` elements ordered by `z-index`
- **`margin-left: auto` right-alignment** for inline-block SVGs in flex containers

### Fixed
- Flex width double-resolution when both `widthUnit` and percentage were set
- Inline-block SVGs disappearing due to missing width propagation

## [0.2.0] - 2026-03-17

### Added
- **Auto-height pages** via CSS `@page { size: 80mm 0; }` — page height sizes to content (receipts, flyers)
- **Negative margin support** for flex column children — enables CSS patterns like `margin: -10px -14px` to break out of parent padding
- **`margin-left: auto`** on flex row items — pushes items to the right edge (e.g., seat box alignment)
- **`margin-top: auto`** on flex column items — pushes items to the bottom (e.g., footer positioning)
- **Cross-axis stretch** (W3C Flexbox §9.4) — flex row items stretch to match tallest sibling, with or without definite container height
- **Flex column `flex-grow`** — items with `flex: 1` now grow to fill remaining space in column direction
- **Flex column `justify-content`** — space-between, center, flex-end, space-around, space-evenly in column direction
- **`hasDefiniteCrossSize` flag** on Flex — enables stretch when Flex is wrapped in a height-constrained Div
- **Watermark support** in WASM render API via `watermark` parameter
- **Automatic Unicode font embedding** for non-WinAnsi characters (CIDFont with embedded cmap)
- **CIDFont fallback decoding** from embedded font cmap tables
- **Font caching** for repeated font resolution
- **Form XObject resolution** in PDF reader
- **Tagged PDF extraction** improvements
- **Full text matrix tracking** and font-aware space detection in reader
- **Xref cycle detection**, hybrid xref support, stream length correction
- **SVG enhancements**: text-anchor, tspan, defs/use, gradient support

### Fixed
- **Percentage heights** now resolve against parent container's explicit height, not the page — fixes vertical bar charts overflowing their containers
- **`box-sizing: border-box`** no longer double-subtracts padding from width/height — only border is subtracted since the Div handles padding internally
- **Double-padding on wrapped flex containers** — when a Flex has CSS width/height, visual properties (padding, borders, margins) are cleared from the Flex and applied only to the wrapper Div
- **`letter-spacing` in width measurement** — `Paragraph.MinWidth()` and `MaxWidth()` now include letter-spacing, preventing flex items from being measured too narrow
- **Floating-point overflow in `margin-top: auto`** — added 0.01pt epsilon tolerance to prevent items from silently overflowing due to float rounding
- **`margin-top: auto` phase consistency** — `neededBelow` calculation now includes `marginBottom` of subsequent items in both phase 1 and phase 3
- **SpaceBefore/SpaceAfter doubling** on flex items — element margins are cleared when FlexItem margins take over (Div, Flex, and Paragraph)
- **Background preserved on wrapped Flex** — `min-height` backgrounds now fill the full height (kept on both Div wrapper and inner Flex)
- **`parseFloat` negative numbers** — CSS parser now correctly handles negative values like `-10px`
- **Flex children splitting** into separate items instead of grouping per HTML child
- **SVG shapes invisible** and text mirrored
- **Sequential elements overlapping** by tracking cumulative Y offset in renderer
- **`<br>` tags in paragraphs** and CSS width as flex-basis
- **WASM binary size** halved by excluding `net/http` from js builds

### Changed
- Cross-axis stretch now fires for all flex row items (not just when container has definite height)
- `planColumn` refactored into 3-phase layout: measure, grow, position

## [0.1.1] - 2026-03-16

### Added
- CLI `extract` command with pluggable strategies (simple, location)
- CLI `sign` command for PAdES digital signatures
- Table `SetBorderCollapse(true)` for CSS-style collapsed borders
- CSS `calc()` support in HTML-to-PDF (e.g., `width: calc(100% - 40px)`)
- CSS `@page` rule parsing (size, margins) in HTML-to-PDF
- CSS `orphans`/`widows` properties in HTML-to-PDF
- CSS `break-before`/`break-after`/`break-inside` modern syntax
- Remote image loading (`<img src="https://...">`) in HTML-to-PDF
- Data URI image support (`<img src="data:image/png;base64,...">`)
- PDF metadata extraction from HTML `<title>` and `<meta>` tags
- Content stream processor with full graphics state (CTM, color, font)
- Pluggable text extraction strategies (Simple, Location, Region)
- Path and image extraction from content streams
- Per-glyph span extraction (opt-in)
- Text rendering mode awareness (invisible text filtering)
- Marked content tag tracking (BMC/BDC/EMC)
- Form XObject recursion in content processing
- Actual glyph widths from font metrics (replaces estimation)
- Auto-bookmarks from layout headings
- Viewer preferences (page layout, mode, UI options)
- Page labels (decimal, Roman, alpha)
- Page geometry boxes (CropBox, BleedBox, TrimBox, ArtBox)
- SVG package in README

### Changed
- CLI version bumped to 0.1.1
- README updated with extract and sign commands, border-collapse, SVG package

### Fixed
- Table border-collapse: adjacent cells no longer draw double borders
- Tables section in README had undefined variable

## [0.1.0] - 2026-03-15

### Added
- Initial release
- PDF generation with layout engine (Paragraph, Heading, Table, List, Div, Image, Float, Flex, Columns)
- PDF reader with tokenizer, parser, xref streams, object streams
- PDF merge and modify
- HTML-to-PDF conversion with CSS support
- Digital signatures (PAdES B-B, B-T, B-LT)
- Interactive forms (AcroForms)
- Barcodes (Code128, QR, EAN-13)
- Tagged PDF and PDF/A compliance
- SVG rendering
- CLI tool (merge, info, pages, text, create, blank)
- Font embedding and subsetting (TrueType)
- JPEG, PNG, TIFF image support
- Encryption (AES-256, AES-128, RC4)
