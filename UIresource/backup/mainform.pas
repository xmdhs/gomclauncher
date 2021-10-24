unit MainForm;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ExtCtrls, StdCtrls,
  ComCtrls;

type

  { TMainForm }

  TMainForm = class(TForm)
    HeadImage: TImage;
    HeadImage1: TImage;
    HeadImage2: TImage;
    HeadImage3: TImage;
    ListBtn: TImage;
    ListBtn1: TImage;
    ListBtn2: TImage;
    GameGroupBox: TGroupBox;
    GeneralSettingBox: TGroupBox;
    ListBtn3: TImage;
    ListView1: TListView;
    T4: TLabel;
    UserName: TLabel;
    UserMode: TLabel;
    UserGroupBox: TGroupBox;
    StaticText1: TStaticText;
    VersionLabel: TLabel;
    T5: TLabel;
    VersionListChose: TComboBox;
    GameStart: TButton;
    GroupBox1: TGroupBox;
    Image1: TImage;
    Panel: TPanel;
    procedure FormShow(Sender: TObject);
    procedure ListBtn1Click(Sender: TObject);
    procedure ListBtn2Click(Sender: TObject);
    procedure ListBtn3Click(Sender: TObject);
    procedure ListBtnClick(Sender: TObject);
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

{ TMainForm }

procedure TMainForm.FormShow(Sender: TObject);
begin

end;

procedure TMainForm.ListBtn1Click(Sender: TObject);
begin

end;

procedure TMainForm.ListBtn2Click(Sender: TObject);
begin

end;

procedure TMainForm.ListBtn3Click(Sender: TObject);
begin

end;

procedure TMainForm.ListBtnClick(Sender: TObject);
begin

end;

end.

