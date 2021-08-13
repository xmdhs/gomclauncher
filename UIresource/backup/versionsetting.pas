unit VersionSetting;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, StdCtrls, ExtCtrls,
  Grids;

type

  { TVersionSetting }

  TVersionSetting = class(TForm)
    Image1: TImage;
    LabeledEdit1: TLabeledEdit;
    ListBox1: TListBox;
    ListBox2: TListBox;
    procedure FormShow(Sender: TObject);
    procedure Image1Click(Sender: TObject);
  private

  public

  end;

var
  VersionSetting: TVersionSetting;

implementation

{$R *.lfm}

{ TVersionSetting }

procedure TVersionSetting.Image1Click(Sender: TObject);
begin

end;

procedure TVersionSetting.FormShow(Sender: TObject);
begin

end;

end.

