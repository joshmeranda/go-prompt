package prompt

type History interface {
	// Add user inputs to the history.
	Add(input ...string)

	// Clear the history to point to the start of the history (the present)
	Clear()

	// Older saves a buffer of current line and get a buffer of previous line by up-arrow.
	// The changes of line buffers are stored until new history is created.
	Older(buf *Buffer) (new *Buffer, changed bool)

	// Newer saves a buffer of current line and get a buffer of next line by down-arrow.
	// The changes of line buffers are stored until new history is created.
	Newer(buf *Buffer) (new *Buffer, changed bool)
}

func NewHistory() History {
	return &history{
		histories: []string{},
		tmp:       []string{""},
		selected:  0,
	}
}

// History stores the texts that are entered.
type history struct {
	histories []string
	tmp       []string
	selected  int
}

// Add to add text in history.
func (h *history) Add(input ...string) {
	h.histories = append(h.histories, input...)
	h.Clear()
}

// Clear to clear the history.
func (h *history) Clear() {
	h.tmp = make([]string, len(h.histories))
	copy(h.tmp, h.histories)
	h.tmp = append(h.tmp, "")
	h.selected = len(h.tmp) - 1
}

func (h *history) Older(buf *Buffer) (new *Buffer, changed bool) {
	if len(h.tmp) == 1 || h.selected == 0 {
		return buf, false
	}
	h.tmp[h.selected] = buf.Text()

	h.selected--
	new = NewBuffer()
	new.InsertText(h.tmp[h.selected], false, true)
	return new, true
}

func (h *history) Newer(buf *Buffer) (new *Buffer, changed bool) {
	if h.selected >= len(h.tmp)-1 {
		return buf, false
	}
	h.tmp[h.selected] = buf.Text()

	h.selected++
	new = NewBuffer()
	new.InsertText(h.tmp[h.selected], false, true)
	return new, true
}
