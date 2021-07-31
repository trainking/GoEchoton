package flatbuf

import (
	sample "GoEchoton/pkg/flatbuf/MyGame/Sample"

	flatbuffers "github.com/google/flatbuffers/go"
)

func NewDemo() {
	builder := flatbuffers.NewBuilder(1024)
	weaponOne := builder.CreateString("Sword")
	weaponTwo := builder.CreateString("Axe")

	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponOne)
	sample.WeaponAddDamage(builder, 3)
	sword := sample.WeaponEnd(builder)

	sample.WeaponStart(builder)
	sample.WeaponAddName(builder, weaponTwo)
	sample.WeaponAddDamage(builder, 5)
	axe := sample.WeaponEnd(builder)

	sample.MonsterStartWeaponsVector(builder, 2)
	builder.PrependUOffsetT(axe)
	builder.PrependUOffsetT(sword)
	weapons := builder.EndVector(2)

	name := builder.CreateString("Orc")

	sample.MonsterStartInventoryVector(builder, 10)
	for i := 9; i >= 0; i-- {
		builder.PrependByte(byte(i))
	}
	inv := builder.EndVector(10)

	sample.MonsterStartPathVector(builder, 2)
	sample.CreateVec3(builder, 1.0, 2.0, 3.0)
	sample.CreateVec3(builder, 4.0, 5.0, 6.0)
	path := builder.EndVector(2)

	sample.MonsterStart(builder)
	sample.MonsterAddPos(builder, sample.CreateVec3(builder, 1.0, 2.0, 3.0))
	sample.MonsterAddHp(builder, 300)
	sample.MonsterAddName(builder, name)
	sample.MonsterAddInventory(builder, inv)
	sample.MonsterAddColor(builder, sample.ColorRed)
	sample.MonsterAddWeapons(builder, weapons)
	sample.MonsterAddEquippedType(builder, sample.EquipmentWeapon)
	sample.MonsterAddEquipped(builder, axe)
	sample.MonsterAddPath(builder, path)
	orc := sample.MonsterEnd(builder)

	builder.Finish(orc)

	buf := builder.FinishedBytes()

	monster := sample.GetRootAsMonster(buf, 0)

	hp := monster.Hp()
	mana := monster.Mana()
	name := string(monster.Name())

	monster.Pos(nil)
}
