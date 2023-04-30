package test

import (
	"image"
	"image/color"
	"image/draw"
	"testing"
	"time"

	"github.com/AlinaDubchak/Lab3-Go/painter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/exp/shiny/screen"
)

type Mock struct {
	mock.Mock
}

func (m *Mock) Update(t screen.Texture) {
	m.Called(t)
}

func (m *Mock) NewBuffer(size image.Point) (screen.Buffer, error) {
	return nil, nil
}

func (m *Mock) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	return nil, nil
}

func (m *Mock) NewTexture(size image.Point) (screen.Texture, error) {
	args := m.Called(size)
	return args.Get(0).(screen.Texture), args.Error(1)
}

func (m *Mock) Release() {
	m.Called()
}

func (m *Mock) Upload(dp image.Point, src screen.Buffer, sr image.Rectangle) {
	m.Called(dp, src, sr)
}

func (m *Mock) Bounds() image.Rectangle {
	args := m.Called()
	return args.Get(0).(image.Rectangle)
}

func (m *Mock) Fill(dr image.Rectangle, src color.Color, op draw.Op) {
	m.Called(dr, src, op)
}

func (m *Mock) Size() image.Point {
	args := m.Called()
	return args.Get(0).(image.Point)
}

func (m *Mock) Do(t screen.Texture) bool {
	args := m.Called(t)
	return args.Bool(0)
}

func TestLoop(t *testing.T) {
	screenMock := new(Mock)
	textureMock := new(Mock)
	receiverMock := new(Mock)
	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	Loop := painter.Loop{
		Receiver: receiverMock,
	}

	Loop.Start(screenMock)

	FirstOperation := new(Mock)
	SecondOperation := new(Mock)

	textureMock.On("Bounds").Return(image.Rectangle{})
	FirstOperation.On("Do", textureMock).Return(false)
	SecondOperation.On("Do", textureMock).Return(true)

	assert.Equal(t, len(Loop.Mq.Queue), 0)
	Loop.Post(FirstOperation)
	Loop.Post(SecondOperation)
	time.Sleep(1 * time.Second)
	assert.Equal(t, len(Loop.Mq.Queue), 0)

	FirstOperation.AssertCalled(t, "Do", textureMock)
	SecondOperation.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}

func TestPass(t *testing.T) {
	screenMock := new(Mock)
	textureMock := new(Mock)
	receiverMock := new(Mock)
	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	Loop := painter.Loop{
		Receiver: receiverMock,
	}

	Loop.Start(screenMock)

	Operation := new(Mock)

	textureMock.On("Bounds").Return(image.Rectangle{})
	Operation.On("Do", textureMock).Return(true)

	assert.Empty(t, len(Loop.Mq.Queue))
	Loop.Post(Operation)
	time.Sleep(1 * time.Second)
	assert.Empty(t, len(Loop.Mq.Queue))

	Operation.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}

func TestFail(t *testing.T) {
	screenMock := new(Mock)
	textureMock := new(Mock)
	receiverMock := new(Mock)
	texture := image.Pt(800, 800)
	screenMock.On("NewTexture", texture).Return(textureMock, nil)
	receiverMock.On("Update", textureMock).Return()
	Loop := painter.Loop{
		Receiver: receiverMock,
	}

	Loop.Start(screenMock)

	Operation := new(Mock)

	textureMock.On("Bounds").Return(image.Rectangle{})
	Operation.On("Do", textureMock).Return(false)

	assert.Empty(t, len(Loop.Mq.Queue))
	Loop.Post(Operation)
	time.Sleep(1 * time.Second)
	assert.Empty(t, len(Loop.Mq.Queue))

	Operation.AssertCalled(t, "Do", textureMock)
	receiverMock.AssertNotCalled(t, "Update", textureMock)
	screenMock.AssertCalled(t, "NewTexture", image.Pt(800, 800))
}